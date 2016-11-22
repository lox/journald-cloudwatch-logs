package main

type Priority int

const (
	EMERGENCY Priority = 0
	ALERT     Priority = 1
	CRITICAL  Priority = 2
	ERROR     Priority = 3
	WARNING   Priority = 4
	NOTICE    Priority = 5
	INFO      Priority = 6
	DEBUG     Priority = 7
)

var PriorityName = map[Priority]string{
	EMERGENCY: "EMERG",
	ALERT:     "ALERT",
	CRITICAL:  "CRITICAL",
	ERROR:     "ERROR",
	WARNING:   "WARNING",
	NOTICE:    "NOTICE",
	INFO:      "INFO",
	DEBUG:     "DEBUG",
}

type Record struct {
	InstanceId  string
	TimeUsec    int64
	PID         int          `journald:"_PID"`
	UID         int          `journald:"_UID"`
	GID         int          `journald:"_GID"`
	Command     string       `journald:"_COMM"`
	Executable  string       `journald:"_EXE"`
	CommandLine string       `journald:"_CMDLINE"`
	SystemdUnit string       `journald:"_SYSTEMD_UNIT"`
	BootId      string       `journald:"_BOOT_ID"`
	MachineId   string       `journald:"_MACHINE_ID"`
	Hostname    string       `journald:"_HOSTNAME"`
	Transport   string       `journald:"_TRANSPORT"`
	Priority    Priority     `journald:"PRIORITY"`
	Message     string       `journald:"MESSAGE"`
	MessageId   string       `journald:"MESSAGE_ID"`
	Errno       int          `journald:"ERRNO"`
	Syslog      RecordSyslog
	Kernel      RecordKernel
}

type RecordSyslog struct {
	Facility   int    `journald:"SYSLOG_FACILITY"`
	Identifier string `journald:"SYSLOG_IDENTIFIER"`
	PID        int    `journald:"SYSLOG_PID"`
}

type RecordKernel struct {
	Device    string `journald:"_KERNEL_DEVICE"`
	Subsystem string `journald:"_KERNEL_SUBSYSTEM"`
	SysName   string `journald:"_UDEV_SYSNAME"`
	DevNode   string `journald:"_UDEV_DEVNODE"`
}
