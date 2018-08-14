package v8

import (
	"testing"

	"io/ioutil"
	"os"

	. "github.com/golang/mock/gomock"
	"github.com/iost-official/Go-IOS-Protocol/core/contract"
	"github.com/iost-official/Go-IOS-Protocol/new_vm/database"
	"github.com/iost-official/Go-IOS-Protocol/new_vm/host"
	"strings"
)

var testDataPath = "./test_data/"

func ReadFile(src string) ([]byte, error) {
	fi, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}
	return fd, nil
}

func Init(t *testing.T) *database.Visitor {
	mc := NewController(t)
	defer mc.Finish()
	db := database.NewMockIMultiValue(mc)
	vi := database.NewVisitor(100, db)
	return vi
}

func MyInit(t *testing.T, conName string, optional ...interface{}) (*VM, *host.Host, *contract.Contract) {
	db := database.NewDatabaseFromPath(testDataPath + conName + ".json")
	vi := database.NewVisitor(100, db)

	ctx := host.NewContext(nil)
	ctx.Set("gas_price", uint64(1))
	var gasLimit = uint64(10000);
	if len(optional) > 0 {
		gasLimit = optional[0].(uint64)
	}
	ctx.Set("gas_limit", gasLimit)
	ctx.Set("contract_name", conName)
	h := &host.Host{Ctx: ctx, DB: vi}

	fd, err := ReadFile(testDataPath + conName + ".js")
	if err != nil {
		t.Fatal("Read file failed: ", err.Error())
		return nil, nil, nil
	}
	rawCode := string(fd)

	code := &contract.Contract{
		ID:   conName,
		Code: rawCode,
	}

	e := NewVM()
	e.Init()
	e.SetJSPath("./v8/libjs/")

	return e, h, code
}

func TestEngine_LoadAndCall(t *testing.T) {
	vi := Init(t)
	ctx := host.NewContext(nil)
	ctx.Set("gas_price", uint64(1))
	ctx.Set("gas_limit", uint64(1000000000))
	ctx.Set("contract_name", "contractName")
	tHost := &host.Host{Ctx: ctx, DB: vi}

	code := &contract.Contract{
		ID: "test.js",
		Code: `
var Contract = function() {
}

	Contract.prototype = {
	fibonacci: function(cycles) {
			if (cycles == 0) return 0;
			if (cycles == 1) return 1;
			return this.fibonacci(cycles - 1) + this.fibonacci(cycles - 2);
		}
	}

	module.exports = Contract
`,
	}

	e := NewVM()
	defer e.Release()
	e.Init()
	e.SetJSPath("./v8/libjs/")

	rs, _, err := e.LoadAndCall(tHost, code, "fibonacci", "12")

	if err != nil {
		t.Fatalf("LoadAndCall run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0] != "144" {
		t.Fatalf("LoadAndCall except 144, got %s\n", rs[0])
	}
}

func TestEngine_bigNumber(t *testing.T) {
	e, host, code := MyInit(t, "bignumber1")
	e.LoadAndCall(host, code, "constructor")
	rs, _, err := e.LoadAndCall(host, code, "getVal")

	if err != nil {
		t.Fatalf("LoadAndCall getVal error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "0.0000000000800029" {
		t.Errorf("LoadAndCall except 0.0000000000800029, got %s\n", rs[0])
	}
}

func TestEngine_Storage(t *testing.T) {
	e, host, code := MyInit(t, "storage1")

	e.LoadAndCall(host, code, "constructor")
	rs, _, err := e.LoadAndCall(host, code, "get", "a", "number")
	if err != nil {
		t.Fatalf("LoadAndCall get run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "1000" {
		t.Fatalf("LoadAndCall except mySetVal, got %s\n", rs[0])
	}

	rtn, cost, err := e.LoadAndCall(host, code, "put", "mySetKey", "mySetVal")
	if err != nil {
		t.Fatalf("LoadAndCall put run error: %v\n", err)
	}
	if len(rtn) != 1 || rtn[0].(string) != "0" {
		t.Fatalf("return of put should be ['0']")
	}
	t.Log(cost)

	rs, _, err = e.LoadAndCall(host, code, "get", "mySetKey", "string")
	if err != nil {
		t.Fatalf("LoadAndCall get run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "mySetVal" {
		t.Fatalf("LoadAndCall except mySetVal, got %s\n", rs[0])
	}

	rs, _, err = e.LoadAndCall(host, code,"getThisNum")
	if err != nil {
		t.Fatalf("LoadAndCall getThisNum run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "99" {
		t.Fatalf("LoadAndCall except 99, got %s\n", rs[0])
	}

	rs, _, err = e.LoadAndCall(host, code,"getThisStr")
	if err != nil {
		t.Fatalf("LoadAndCall getThisStr run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "yeah" {
		t.Fatalf("LoadAndCall except yeah, got %s\n", rs[0])
	}

	rtn, cost, err = e.LoadAndCall(host, code, "delete", "mySetKey")
	if err != nil {
		t.Fatalf("LoadAndCall delete run error: %v\n", err)
	}
	if len(rtn) != 1 || rtn[0].(string) != "0" {
		t.Fatalf("return of put should be ['0']")
	}
	t.Log(cost)

	rs, _, err = e.LoadAndCall(host, code, "get", "mySetKey")
	if err != nil {
		t.Fatalf("LoadAndCall get run error: %v\n", err)
	}
	// todo get return nil
	if len(rs) != 1 || rs[0].(string) != "nil" {
		t.Fatalf("LoadAndCall except mySetVal, got %s\n", rs[0])
	}
}

func TestEngine_DataType(t *testing.T) {
	e, host, code := MyInit(t, "datatype")

	rs, _,err := e.LoadAndCall(host, code, "number", 1)
	if err != nil {
		t.Fatalf("LoadAndCall number run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "0.5555555555" {
		t.Fatalf("LoadAndCall except 0.5555555555, got %s\n", rs[0])
	}

	rs, _,err = e.LoadAndCall(host, code, "number_big", 1)
	if err != nil {
		t.Fatalf("LoadAndCall number_big run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "0.5555555555" {
		t.Fatalf("LoadAndCall except 0.5555555555, got %s\n", rs[0])
	}

	rs, _,err = e.LoadAndCall(host, code, "number_op")
	if err != nil {
		t.Fatalf("LoadAndCall number_op run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "3" {
		t.Fatalf("LoadAndCall except 3, got %s\n", rs[0])
	}

	rs, _,err = e.LoadAndCall(host, code, "number_op2")
	if err != nil {
		t.Fatalf("LoadAndCall number_op2 run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "2" {
		t.Fatalf("LoadAndCall except 3, got %s\n", rs[0])
	}

	rs, _,err = e.LoadAndCall(host, code, "number_strange")
	if err != nil {
		t.Fatalf("LoadAndCall number_strange run error: %v\n", err)
	}
	// todo get return string -infinity
	if len(rs) != 1 || rs[0].(string) != "-Infinity" {
		t.Fatalf("LoadAndCall except Infinity, got %s\n", rs[0])
	}

	rs, _,err = e.LoadAndCall(host, code, "param", 2)
	if err != nil {
		t.Fatalf("LoadAndCall param run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "4" {
		t.Fatalf("LoadAndCall except 4, got %s %s\n", rs[0])
	}

	rs, _,err = e.LoadAndCall(host, code, "param2")
	if err != nil {
		t.Fatalf("LoadAndCall param run error: %v\n", err)
	}
	// todo get return string undefined
	if len(rs) != 1 || rs[0].(string) != "undefined" {
		t.Fatalf("LoadAndCall except undefined, got %s\n", rs[0])
	}

	rs, _,err = e.LoadAndCall(host, code, "bool")
	if err != nil {
		t.Fatalf("LoadAndCall bool run error: %v\n", err)
	}
	// todo get return string false
	if len(rs) != 1 || rs[0].(string) != "false" {
		t.Fatalf("LoadAndCall except undefined, got %s\n", rs[0])
	}

	rs, _,err = e.LoadAndCall(host, code, "string")
	if err != nil {
		t.Fatalf("LoadAndCall string run error: %v\n", err)
	}
	if len(rs) != 1 || len(rs[0].(string)) != 4096 {
		t.Fatalf("LoadAndCall except len 4096, got %d\n", len(rs[0].(string)))
	}

	rs, _,err = e.LoadAndCall(host, code, "array")
	if err != nil {
		t.Fatalf("LoadAndCall array run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "0,1,2,3"{
		t.Fatalf("LoadAndCall except 0,1,2,3, got %s\n", rs[0].(string))
	}

	// todo get return string [object Object]
	/*
	rs, _,err = e.LoadAndCall(host, code, "object")
	if err != nil {
		t.Fatalf("LoadAndCall array run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "0,1,2,3"{
		t.Errorf("LoadAndCall except 0,1,2,3, got %s\n", rs[0].(string))
	}
	*/
}

func TestEngine_Loop(t *testing.T) {
	e, host, code := MyInit(t, "loop")

	_, _,err := e.LoadAndCall(host, code, "for")
	if err == nil || err.Error() != "out of gas"{
		t.Fatalf("LoadAndCall for should return error: out of gas, but got %v\n", err)
	}

	e, host, code = MyInit(t, "loop")
	_, _,err = e.LoadAndCall(host, code, "for2")
	if err == nil || err.Error() != "out of gas"{
		t.Fatalf("LoadAndCall for should return error: out of gas, but got %v\n", err)
	}

	e, host, code = MyInit(t, "loop")
	rs, _,err := e.LoadAndCall(host, code, "for3")
	if err != nil {
		t.Fatalf("LoadAndCall for3 run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "10"{
		t.Fatalf("LoadAndCall except 10, got %s\n", rs[0].(string))
	}

	e, host, code = MyInit(t, "loop")
	rs, _,err = e.LoadAndCall(host, code, "forin")
	if err != nil {
		t.Fatalf("LoadAndCall forin run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "12"{
		t.Fatalf("LoadAndCall except 12, got %s\n", rs[0].(string))
	}

	e, host, code = MyInit(t, "loop")
	rs, _,err = e.LoadAndCall(host, code, "forof")
	if err != nil {
		t.Fatalf("LoadAndCall forof run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "6"{
		t.Fatalf("LoadAndCall except 6, got %s\n", rs[0].(string))
	}

	e, host, code = MyInit(t, "loop")
	_, _,err = e.LoadAndCall(host, code, "while")
	if err == nil || err.Error() != "out of gas"{
		t.Fatalf("LoadAndCall for should return error: out of gas, but got %v\n", err)
	}

	e, host, code = MyInit(t, "loop")
	_, _,err = e.LoadAndCall(host, code, "dowhile")
	if err == nil || err.Error() != "out of gas"{
		t.Fatalf("LoadAndCall for should return error: out of gas, but got %v\n", err)
	}
}

func TestEngine_Func(t *testing.T) {
	e, host, code := MyInit(t, "func")
	_, _,err := e.LoadAndCall(host, code, "func1")
	if err == nil || err.Error() != "out of gas"{
		t.Fatalf("LoadAndCall for should return error: out of gas, but got %v\n", err)
	}

	e, host, code = MyInit(t, "func", uint64(100000000000))
	_, _,err = e.LoadAndCall(host, code, "func1")
	if err == nil || !strings.Contains(err.Error(), "Uncaught exception: RangeError: Maximum call stack size exceeded") {
		t.Fatalf("LoadAndCall for should return error: Uncaught exception: RangeError: Maximum call stack size exceeded, but got %v\n", err)
	}

	e, host, code = MyInit(t, "func")
	rs, _,err := e.LoadAndCall(host, code, "func3", 4)
	if err != nil {
		t.Fatalf("LoadAndCall func3 run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "9"{
		t.Fatalf("LoadAndCall except 9, got %s\n", rs[0].(string))
	}

	e, host, code = MyInit(t, "func")
	rs, _,err = e.LoadAndCall(host, code, "func4")
	if err != nil {
		t.Fatalf("LoadAndCall func4 run error: %v\n", err)
	}
	if len(rs) != 1 || rs[0].(string) != "2,5,5"{
		t.Fatalf("LoadAndCall except 2,5,5, got %s\n", rs[0].(string))
	}
}

func TestEngine_Danger(t *testing.T) {
	e, host, code := MyInit(t, "danger")
	_, _,err := e.LoadAndCall(host, code, "bigArray")
	if err == nil || err.Error() != "execution killed"{
		t.Fatalf("LoadAndCall for should return error: execution killed, but got %v\n", err)
	}

	e, host, code = MyInit(t, "danger")
	_, _,err = e.LoadAndCall(host, code, "visitUndefined")
	if err == nil || !strings.Contains(err.Error(),"Uncaught exception: TypeError: Cannot set property 'c' of undefined") {
		t.Fatalf("LoadAndCall for should return error: Uncaught exception: TypeError: Cannot set property 'c' of undefined, but got %v\n", err)
	}

	e, host, code = MyInit(t, "danger")
	_, _,err = e.LoadAndCall(host, code, "throw")
	if err == nil || !strings.Contains(err.Error(),"Uncaught exception: test throw") {
		t.Fatalf("LoadAndCall for should return error: Uncaught exception: test throw, but got %v\n", err)
	}
}
