# Variable for filename for store running procees id
PID_FILE = gopush.pid
DLV_FILE = dlv.pid

# Start task run app and writes it's process id to PID_FILE.
start:
	go build -gcflags="all=-N -l"
	./gopush & echo $$! > $(PID_FILE)
	dlv attach --listen=:40000 --headless --api-version=2 --accept-multiclient --continue `cat $(PID_FILE)` & echo $$! > $(DLV_FILE)

# Stop task will kill process by ID stored in PID_FILE (and all child processes).
stop:
#	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`
#	-pkill -P `cat $(PID_FILE)`
	-kill -15 `cat $(DLV_FILE)`
	-kill -15 `cat $(PID_FILE)`

# Before task will only prints message. Actually, it is not necessary.
before:
	@echo "STOPPED gopush"

# Restart task will execute stop, before and start tasks in strict order and prints message.
restart: stop before start
	@echo "STARTED gopush" && printf '%*s\n' "40" '' | tr ' ' -

# Serve task will run fswatch monitor and performs restart task if source file changed. Before serving it will execute start task.
serve: start
	fswatch -or --event=Updated ./config/config.yml | \
	xargs -n1 -I {} make restart

# .PHONY is used for reserving tasks words
.PHONY: start before stop restart serve

