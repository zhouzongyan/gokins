package cmd

import (
	"github.com/gokins-main/core"
	"github.com/gokins-main/gokins/comm"
	"github.com/gokins-main/gokins/server"
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/exec"
)

const Version = "0.1.1"

var (
	app = kingpin.New("gokins", "A golang workflow application.")
)

func Run() {
	regs()
	kingpin.Version(Version)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
func regs() {
	app.Flag("web", "gokins web host").Default(":8030").StringVar(&comm.WebHost)
	//app.Flag("hbtp", "gokins hbtp host").Default(":8031").StringVar(&comm.HbtpHost)
	app.Flag("workdir", "gokins work path").Short('w').StringVar(&comm.WorkPath)
	cmd := app.Command("run", "run process").Default().
		Action(run)
	cmd.Flag("debug", "debug log show").BoolVar(&core.Debug)

	cmd = app.Command("daemon", "run process background").
		Action(start)
}
func getArgs() []string {
	args := make([]string, 0)
	args = append(args, "run")
	if comm.WebHost != "" {
		args = append(args, "--web")
		args = append(args, comm.WebHost)
	}
	/*if comm.HbtpHost != "" {
		args = append(args, "--hbtp")
		args = append(args, comm.HbtpHost)
	}*/
	if comm.WorkPath != "" {
		args = append(args, "--workdir")
		args = append(args, comm.WorkPath)
	}
	return args
}
func start(pc *kingpin.ParseContext) error {
	args := getArgs()
	fullpth, err := os.Executable()
	if err != nil {
		return err
	}
	println("start process")
	cmd := exec.Command(fullpth, args...)
	err = cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
func run(pc *kingpin.ParseContext) error {
	if core.Debug {
		hbtp.Debug = true
	}
	return server.Run()
	/*err:=server.RunHbtp()
	if err!=nil{
		return err
	}
	mgr := monitor.NewManager(DevName)
	monitor.YwtEgn = ywtree.NewEngine(mgr, &ywtree.Config{
		Host:   HbtpHost,
		Org:    "mgr",
		Name:   "cpu-monitor",
		Secret: DevSecret,
	})
	monitor.YwtEgn.RegGrpcGrpcFun(1, &monitor.DevRPC{})
	return monitor.YwtEgn.Run()*/
}