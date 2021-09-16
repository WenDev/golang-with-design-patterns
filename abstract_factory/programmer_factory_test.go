package abstract_factory

import (
	"fmt"
	"testing"
)

func TestCreateProgrammerAndArchitect(t *testing.T) {
	fmt.Println("前端组招到了一个程序员和一个架构师")
	ff := FrontEndFactory{}
	fa := ff.CreateArchitect()
	fp := ff.CreateProgrammer()
	fmt.Println("前端组接到任务,开始工作...")
	fa.Design()
	fp.Work()
	fmt.Println("后端组招到了一个程序员和一个架构师")
	bf := BackEndFactory{}
	ba := bf.CreateArchitect()
	bp := bf.CreateProgrammer()
	fmt.Println("后端组接到任务,开始工作...")
	ba.Design()
	bp.Work()
}
