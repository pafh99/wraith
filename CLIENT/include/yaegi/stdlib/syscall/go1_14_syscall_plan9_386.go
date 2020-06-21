// Code generated by 'goexports syscall'. DO NOT EDIT.

// +build go1.14,!go1.15

package syscall

import (
	"go/constant"
	"go/token"
	"reflect"
	"syscall"
)

func init() {
	Symbols["syscall"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Await":               reflect.ValueOf(syscall.Await),
		"Bind":                reflect.ValueOf(syscall.Bind),
		"BytePtrFromString":   reflect.ValueOf(syscall.BytePtrFromString),
		"ByteSliceFromString": reflect.ValueOf(syscall.ByteSliceFromString),
		"Chdir":               reflect.ValueOf(syscall.Chdir),
		"Clearenv":            reflect.ValueOf(syscall.Clearenv),
		"Close":               reflect.ValueOf(syscall.Close),
		"Create":              reflect.ValueOf(syscall.Create),
		"DMAPPEND":            reflect.ValueOf(constant.MakeFromLiteral("1073741824", token.INT, 0)),
		"DMAUTH":              reflect.ValueOf(constant.MakeFromLiteral("134217728", token.INT, 0)),
		"DMDIR":               reflect.ValueOf(constant.MakeFromLiteral("2147483648", token.INT, 0)),
		"DMEXCL":              reflect.ValueOf(constant.MakeFromLiteral("536870912", token.INT, 0)),
		"DMEXEC":              reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"DMMOUNT":             reflect.ValueOf(constant.MakeFromLiteral("268435456", token.INT, 0)),
		"DMREAD":              reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"DMTMP":               reflect.ValueOf(constant.MakeFromLiteral("67108864", token.INT, 0)),
		"DMWRITE":             reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"Dup":                 reflect.ValueOf(syscall.Dup),
		"EACCES":              reflect.ValueOf(&syscall.EACCES).Elem(),
		"EAFNOSUPPORT":        reflect.ValueOf(&syscall.EAFNOSUPPORT).Elem(),
		"EBUSY":               reflect.ValueOf(&syscall.EBUSY).Elem(),
		"EEXIST":              reflect.ValueOf(&syscall.EEXIST).Elem(),
		"EINTR":               reflect.ValueOf(&syscall.EINTR).Elem(),
		"EINVAL":              reflect.ValueOf(&syscall.EINVAL).Elem(),
		"EIO":                 reflect.ValueOf(&syscall.EIO).Elem(),
		"EISDIR":              reflect.ValueOf(&syscall.EISDIR).Elem(),
		"EMFILE":              reflect.ValueOf(&syscall.EMFILE).Elem(),
		"ENAMETOOLONG":        reflect.ValueOf(&syscall.ENAMETOOLONG).Elem(),
		"ENOENT":              reflect.ValueOf(&syscall.ENOENT).Elem(),
		"ENOTDIR":             reflect.ValueOf(&syscall.ENOTDIR).Elem(),
		"EPERM":               reflect.ValueOf(&syscall.EPERM).Elem(),
		"EPLAN9":              reflect.ValueOf(&syscall.EPLAN9).Elem(),
		"ERRMAX":              reflect.ValueOf(constant.MakeFromLiteral("128", token.INT, 0)),
		"ESPIPE":              reflect.ValueOf(&syscall.ESPIPE).Elem(),
		"ETIMEDOUT":           reflect.ValueOf(&syscall.ETIMEDOUT).Elem(),
		"Environ":             reflect.ValueOf(syscall.Environ),
		"ErrBadName":          reflect.ValueOf(&syscall.ErrBadName).Elem(),
		"ErrBadStat":          reflect.ValueOf(&syscall.ErrBadStat).Elem(),
		"ErrShortStat":        reflect.ValueOf(&syscall.ErrShortStat).Elem(),
		"Exec":                reflect.ValueOf(syscall.Exec),
		"Exit":                reflect.ValueOf(syscall.Exit),
		"Fchdir":              reflect.ValueOf(syscall.Fchdir),
		"Fd2path":             reflect.ValueOf(syscall.Fd2path),
		"Fixwd":               reflect.ValueOf(syscall.Fixwd),
		"ForkExec":            reflect.ValueOf(syscall.ForkExec),
		"ForkLock":            reflect.ValueOf(&syscall.ForkLock).Elem(),
		"Fstat":               reflect.ValueOf(syscall.Fstat),
		"Fwstat":              reflect.ValueOf(syscall.Fwstat),
		"Getegid":             reflect.ValueOf(syscall.Getegid),
		"Getenv":              reflect.ValueOf(syscall.Getenv),
		"Geteuid":             reflect.ValueOf(syscall.Geteuid),
		"Getgid":              reflect.ValueOf(syscall.Getgid),
		"Getgroups":           reflect.ValueOf(syscall.Getgroups),
		"Getpagesize":         reflect.ValueOf(syscall.Getpagesize),
		"Getpid":              reflect.ValueOf(syscall.Getpid),
		"Getppid":             reflect.ValueOf(syscall.Getppid),
		"Gettimeofday":        reflect.ValueOf(syscall.Gettimeofday),
		"Getuid":              reflect.ValueOf(syscall.Getuid),
		"Getwd":               reflect.ValueOf(syscall.Getwd),
		"ImplementsGetwd":     reflect.ValueOf(syscall.ImplementsGetwd),
		"MAFTER":              reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"MBEFORE":             reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"MCACHE":              reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"MCREATE":             reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"MMASK":               reflect.ValueOf(constant.MakeFromLiteral("23", token.INT, 0)),
		"MORDER":              reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"MREPL":               reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"Mkdir":               reflect.ValueOf(syscall.Mkdir),
		"Mount":               reflect.ValueOf(syscall.Mount),
		"NewError":            reflect.ValueOf(syscall.NewError),
		"NsecToTimeval":       reflect.ValueOf(syscall.NsecToTimeval),
		"O_APPEND":            reflect.ValueOf(constant.MakeFromLiteral("1024", token.INT, 0)),
		"O_ASYNC":             reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"O_CLOEXEC":           reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"O_CREAT":             reflect.ValueOf(constant.MakeFromLiteral("8192", token.INT, 0)),
		"O_EXCL":              reflect.ValueOf(constant.MakeFromLiteral("4096", token.INT, 0)),
		"O_NOCTTY":            reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"O_NONBLOCK":          reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"O_RDONLY":            reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"O_RDWR":              reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"O_SYNC":              reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"O_TRUNC":             reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"O_WRONLY":            reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"Open":                reflect.ValueOf(syscall.Open),
		"Pipe":                reflect.ValueOf(syscall.Pipe),
		"Pread":               reflect.ValueOf(syscall.Pread),
		"Pwrite":              reflect.ValueOf(syscall.Pwrite),
		"QTAPPEND":            reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"QTAUTH":              reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"QTDIR":               reflect.ValueOf(constant.MakeFromLiteral("128", token.INT, 0)),
		"QTEXCL":              reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"QTFILE":              reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"QTMOUNT":             reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"QTTMP":               reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"RFCENVG":             reflect.ValueOf(constant.MakeFromLiteral("2048", token.INT, 0)),
		"RFCFDG":              reflect.ValueOf(constant.MakeFromLiteral("4096", token.INT, 0)),
		"RFCNAMEG":            reflect.ValueOf(constant.MakeFromLiteral("1024", token.INT, 0)),
		"RFENVG":              reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"RFFDG":               reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"RFMEM":               reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"RFNAMEG":             reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"RFNOMNT":             reflect.ValueOf(constant.MakeFromLiteral("16384", token.INT, 0)),
		"RFNOTEG":             reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"RFNOWAIT":            reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"RFPROC":              reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"RFREND":              reflect.ValueOf(constant.MakeFromLiteral("8192", token.INT, 0)),
		"RawSyscall":          reflect.ValueOf(syscall.RawSyscall),
		"RawSyscall6":         reflect.ValueOf(syscall.RawSyscall6),
		"Read":                reflect.ValueOf(syscall.Read),
		"Remove":              reflect.ValueOf(syscall.Remove),
		"SIGABRT":             reflect.ValueOf(syscall.SIGABRT),
		"SIGALRM":             reflect.ValueOf(syscall.SIGALRM),
		"SIGHUP":              reflect.ValueOf(syscall.SIGHUP),
		"SIGINT":              reflect.ValueOf(syscall.SIGINT),
		"SIGKILL":             reflect.ValueOf(syscall.SIGKILL),
		"SIGTERM":             reflect.ValueOf(syscall.SIGTERM),
		"STATFIXLEN":          reflect.ValueOf(constant.MakeFromLiteral("49", token.INT, 0)),
		"STATMAX":             reflect.ValueOf(constant.MakeFromLiteral("65535", token.INT, 0)),
		"SYS_ALARM":           reflect.ValueOf(constant.MakeFromLiteral("6", token.INT, 0)),
		"SYS_AWAIT":           reflect.ValueOf(constant.MakeFromLiteral("47", token.INT, 0)),
		"SYS_BIND":            reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"SYS_BRK_":            reflect.ValueOf(constant.MakeFromLiteral("24", token.INT, 0)),
		"SYS_CHDIR":           reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"SYS_CLOSE":           reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"SYS_CREATE":          reflect.ValueOf(constant.MakeFromLiteral("22", token.INT, 0)),
		"SYS_DUP":             reflect.ValueOf(constant.MakeFromLiteral("5", token.INT, 0)),
		"SYS_ERRSTR":          reflect.ValueOf(constant.MakeFromLiteral("41", token.INT, 0)),
		"SYS_EXEC":            reflect.ValueOf(constant.MakeFromLiteral("7", token.INT, 0)),
		"SYS_EXITS":           reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"SYS_FAUTH":           reflect.ValueOf(constant.MakeFromLiteral("10", token.INT, 0)),
		"SYS_FD2PATH":         reflect.ValueOf(constant.MakeFromLiteral("23", token.INT, 0)),
		"SYS_FSTAT":           reflect.ValueOf(constant.MakeFromLiteral("43", token.INT, 0)),
		"SYS_FVERSION":        reflect.ValueOf(constant.MakeFromLiteral("40", token.INT, 0)),
		"SYS_FWSTAT":          reflect.ValueOf(constant.MakeFromLiteral("45", token.INT, 0)),
		"SYS_MOUNT":           reflect.ValueOf(constant.MakeFromLiteral("46", token.INT, 0)),
		"SYS_NOTED":           reflect.ValueOf(constant.MakeFromLiteral("29", token.INT, 0)),
		"SYS_NOTIFY":          reflect.ValueOf(constant.MakeFromLiteral("28", token.INT, 0)),
		"SYS_NSEC":            reflect.ValueOf(constant.MakeFromLiteral("53", token.INT, 0)),
		"SYS_OPEN":            reflect.ValueOf(constant.MakeFromLiteral("14", token.INT, 0)),
		"SYS_OSEEK":           reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"SYS_PIPE":            reflect.ValueOf(constant.MakeFromLiteral("21", token.INT, 0)),
		"SYS_PREAD":           reflect.ValueOf(constant.MakeFromLiteral("50", token.INT, 0)),
		"SYS_PWRITE":          reflect.ValueOf(constant.MakeFromLiteral("51", token.INT, 0)),
		"SYS_REMOVE":          reflect.ValueOf(constant.MakeFromLiteral("25", token.INT, 0)),
		"SYS_RENDEZVOUS":      reflect.ValueOf(constant.MakeFromLiteral("34", token.INT, 0)),
		"SYS_RFORK":           reflect.ValueOf(constant.MakeFromLiteral("19", token.INT, 0)),
		"SYS_SEEK":            reflect.ValueOf(constant.MakeFromLiteral("39", token.INT, 0)),
		"SYS_SEGATTACH":       reflect.ValueOf(constant.MakeFromLiteral("30", token.INT, 0)),
		"SYS_SEGBRK":          reflect.ValueOf(constant.MakeFromLiteral("12", token.INT, 0)),
		"SYS_SEGDETACH":       reflect.ValueOf(constant.MakeFromLiteral("31", token.INT, 0)),
		"SYS_SEGFLUSH":        reflect.ValueOf(constant.MakeFromLiteral("33", token.INT, 0)),
		"SYS_SEGFREE":         reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"SYS_SEMACQUIRE":      reflect.ValueOf(constant.MakeFromLiteral("37", token.INT, 0)),
		"SYS_SEMRELEASE":      reflect.ValueOf(constant.MakeFromLiteral("38", token.INT, 0)),
		"SYS_SLEEP":           reflect.ValueOf(constant.MakeFromLiteral("17", token.INT, 0)),
		"SYS_STAT":            reflect.ValueOf(constant.MakeFromLiteral("42", token.INT, 0)),
		"SYS_SYSR1":           reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"SYS_TSEMACQUIRE":     reflect.ValueOf(constant.MakeFromLiteral("52", token.INT, 0)),
		"SYS_UNMOUNT":         reflect.ValueOf(constant.MakeFromLiteral("35", token.INT, 0)),
		"SYS_WSTAT":           reflect.ValueOf(constant.MakeFromLiteral("44", token.INT, 0)),
		"S_IFBLK":             reflect.ValueOf(constant.MakeFromLiteral("24576", token.INT, 0)),
		"S_IFCHR":             reflect.ValueOf(constant.MakeFromLiteral("8192", token.INT, 0)),
		"S_IFDIR":             reflect.ValueOf(constant.MakeFromLiteral("16384", token.INT, 0)),
		"S_IFIFO":             reflect.ValueOf(constant.MakeFromLiteral("4096", token.INT, 0)),
		"S_IFLNK":             reflect.ValueOf(constant.MakeFromLiteral("40960", token.INT, 0)),
		"S_IFMT":              reflect.ValueOf(constant.MakeFromLiteral("126976", token.INT, 0)),
		"S_IFREG":             reflect.ValueOf(constant.MakeFromLiteral("32768", token.INT, 0)),
		"S_IFSOCK":            reflect.ValueOf(constant.MakeFromLiteral("49152", token.INT, 0)),
		"Seek":                reflect.ValueOf(syscall.Seek),
		"Setenv":              reflect.ValueOf(syscall.Setenv),
		"SlicePtrFromStrings": reflect.ValueOf(syscall.SlicePtrFromStrings),
		"SocketDisableIPv6":   reflect.ValueOf(&syscall.SocketDisableIPv6).Elem(),
		"StartProcess":        reflect.ValueOf(syscall.StartProcess),
		"Stat":                reflect.ValueOf(syscall.Stat),
		"Stderr":              reflect.ValueOf(&syscall.Stderr).Elem(),
		"Stdin":               reflect.ValueOf(&syscall.Stdin).Elem(),
		"Stdout":              reflect.ValueOf(&syscall.Stdout).Elem(),
		"StringBytePtr":       reflect.ValueOf(syscall.StringBytePtr),
		"StringByteSlice":     reflect.ValueOf(syscall.StringByteSlice),
		"StringSlicePtr":      reflect.ValueOf(syscall.StringSlicePtr),
		"Syscall":             reflect.ValueOf(syscall.Syscall),
		"Syscall6":            reflect.ValueOf(syscall.Syscall6),
		"UnmarshalDir":        reflect.ValueOf(syscall.UnmarshalDir),
		"Unmount":             reflect.ValueOf(syscall.Unmount),
		"Unsetenv":            reflect.ValueOf(syscall.Unsetenv),
		"WaitProcess":         reflect.ValueOf(syscall.WaitProcess),
		"Write":               reflect.ValueOf(syscall.Write),
		"Wstat":               reflect.ValueOf(syscall.Wstat),

		// type definitions
		"Conn":        reflect.ValueOf((*syscall.Conn)(nil)),
		"Dir":         reflect.ValueOf((*syscall.Dir)(nil)),
		"ErrorString": reflect.ValueOf((*syscall.ErrorString)(nil)),
		"Note":        reflect.ValueOf((*syscall.Note)(nil)),
		"ProcAttr":    reflect.ValueOf((*syscall.ProcAttr)(nil)),
		"Qid":         reflect.ValueOf((*syscall.Qid)(nil)),
		"RawConn":     reflect.ValueOf((*syscall.RawConn)(nil)),
		"SysProcAttr": reflect.ValueOf((*syscall.SysProcAttr)(nil)),
		"Timespec":    reflect.ValueOf((*syscall.Timespec)(nil)),
		"Timeval":     reflect.ValueOf((*syscall.Timeval)(nil)),
		"Waitmsg":     reflect.ValueOf((*syscall.Waitmsg)(nil)),

		// interface wrapper definitions
		"_Conn":    reflect.ValueOf((*_syscall_Conn)(nil)),
		"_RawConn": reflect.ValueOf((*_syscall_RawConn)(nil)),
	}
}

// _syscall_Conn is an interface wrapper for Conn type
type _syscall_Conn struct {
	WSyscallConn func() (syscall.RawConn, error)
}

func (W _syscall_Conn) SyscallConn() (syscall.RawConn, error) { return W.WSyscallConn() }

// _syscall_RawConn is an interface wrapper for RawConn type
type _syscall_RawConn struct {
	WControl func(f func(fd uintptr)) error
	WRead    func(f func(fd uintptr) (done bool)) error
	WWrite   func(f func(fd uintptr) (done bool)) error
}

func (W _syscall_RawConn) Control(f func(fd uintptr)) error           { return W.WControl(f) }
func (W _syscall_RawConn) Read(f func(fd uintptr) (done bool)) error  { return W.WRead(f) }
func (W _syscall_RawConn) Write(f func(fd uintptr) (done bool)) error { return W.WWrite(f) }
