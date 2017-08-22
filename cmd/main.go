package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/krzysiekpiasecki/burzedzisnet"
)

// persistent auth with header?
// panic when timeout handle it
// apikey strategies and caching strategies
// Round position x and y to precision 2
// After authorization set key to enviroment or cache it?

type cmd struct {
	name    string
	flagSet *flag.FlagSet
}

func (c cmd) parse(args []string) error {
	return c.flagSet.Parse(args)
}

type authCmdArgs struct {
	apikey string
}

type locateCmdArgs struct {
	apikey string
	name   string
}

type MyComplexTypeOstrzezeniaCmdArgs struct {
	apikey string
	name   string
	x      float64
	y      float64
}

type MyComplexTypeBurzaCmdArgs struct {
	apikey   string
	location string
	x        float64
	y        float64
	radius   int
}

func authCmd(args *authCmdArgs) cmd {
	fs := flag.NewFlagSet("auth", flag.ExitOnError)
	fs.StringVar(&(*args).apikey, "apikey", "", "Personal key to access a burzedzisnet API")
	c := cmd{"auth", fs}
	return c
}

func locateCmd(args *locateCmdArgs) cmd {
	fs := flag.NewFlagSet("locate", flag.ExitOnError)
	fs.StringVar(&(*args).apikey, "apikey", "", "Personal key to access a burzedzisnet API")
	fs.StringVar(&(*args).name, "name", "", "Location name")
	c := cmd{"locate", fs}
	return c
}

func MyComplexTypeOstrzezeniaCmd(args *MyComplexTypeOstrzezeniaCmdArgs) cmd {
	fs := flag.NewFlagSet("myComplexTypeOstrzezenia", flag.ExitOnError)
	fs.StringVar(&(*args).apikey, "apikey", "", "Personal key to access a burzedzisnet API")
	fs.StringVar(&(*args).name, "name", "", "Location name")
	fs.Float64Var(&(*args).x, "x", 0, "myComplexTypeMiejscowosc x of myComplexTypeMiejscowosc")
	fs.Float64Var(&(*args).y, "y", 0, "myComplexTypeMiejscowosc y of myComplexTypeMiejscowosc")
	c := cmd{"MyComplexTypeOstrzezenias", fs}
	return c
}

func MyComplexTypeBurzaCmd(args *MyComplexTypeBurzaCmdArgs) cmd {
	fs := flag.NewFlagSet("myComplexTypeBurza", flag.ExitOnError)
	fs.StringVar(&(*args).apikey, "apikey", "", "Personal key to access a burzedzisnet API")
	fs.StringVar(&(*args).location, "location", "", "Location name")
	fs.Float64Var(&(*args).x, "x", 0, "myComplexTypeMiejscowosc x of myComplexTypeMiejscowosc")
	fs.Float64Var(&(*args).y, "y", 0, "myComplexTypeMiejscowosc y of myComplexTypeMiejscowosc")
	fs.IntVar(&(*args).radius, "radius", 0, "The radius covered by point, optional. Default is 25 km.")
	c := cmd{"search", fs}
	return c
}

func (args authCmdArgs) print(m *map[string]string) *map[string]string {
	(*m)["apikey"] = args.apikey
	return m
}

func (args locateCmdArgs) print(m *map[string]string) *map[string]string {
	(*m)["apikey"] = args.apikey
	(*m)["name"] = args.name
	return m
}

func (args MyComplexTypeOstrzezeniaCmdArgs) print(m *map[string]string) *map[string]string {
	(*m)["apikey"] = args.apikey
	(*m)["x"] = strconv.FormatFloat(args.x, 'f', 2, 64)
	(*m)["y"] = strconv.FormatFloat(args.y, 'f', 2, 64)
	(*m)["name"] = args.name
	return m
}

func (args MyComplexTypeBurzaCmdArgs) print(m *map[string]string) *map[string]string {
	(*m)["apikey"] = args.apikey
	(*m)["x"] = strconv.FormatFloat(args.x, 'f', 2, 64)
	(*m)["y"] = strconv.FormatFloat(args.y, 'f', 2, 64)
	(*m)["location"] = args.location
	(*m)["radius"] = strconv.FormatInt(int64(args.radius), 10)
	return m
}

func main() {

	if len(os.Args) == 1 {
		fmt.Print("Provide location name")
		os.Exit(1)
	}

	fs := flag.NewFlagSet("bdn", flag.ExitOnError)
	apiKey := fs.String("apikey", "default", "Api key")
	fs.Parse(os.Args[2:])

	client := burzedzisnet.NewClient(*apiKey)

	loc, _ := client.MyComplexTypeMiejscowosc(os.Args[1])
	MyComplexTypeBurza, _ := client.MyComplexTypeBurza(loc.X, loc.Y, 25)
	warn, _ := client.MyComplexTypeOstrzezenia(loc.X, loc.Y)
	fmt.Println(loc)
	fmt.Println(MyComplexTypeBurza)

	if warn.IsSafe() == false {
		fmt.Println(warn)
	} else {
		fmt.Println("No weather alerts!")
	}

}
