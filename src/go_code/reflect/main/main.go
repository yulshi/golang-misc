package main

import (
  "fmt"
  "go_code/reflect/model"
  "reflect"
)

func main() {

  //student := model.Student{
  //  Name: "Tom",
  //  //Age:   20,
  //  Score: 78.5,
  //}
  //reflectStruct(&student)

  //add := func(n1, n2 int) int {
  // return n1 + n2
  //}
  ////print := func() {
  ////  println("Hello Go...")
  ////}
  //adapter(add , 4, 6)

  dynamicCreate()

}

func dynamicCreate() {

  var student *model.Student

  // 根据变量定义获取type，并得到里面的实际类型
  refType := reflect.TypeOf(student).Elem()
  // 根据类型新建一个Value，类型是Ptr
  refValue := reflect.New(refType)

  student = refValue.Interface().(*model.Student)

  // 获取Ptr指向的实际值
  elem := refValue.Elem()

  // 为每个字段赋值
  elem.FieldByName("Name").SetString("Tom")
  elem.FieldByName("Age").SetInt(20)
  elem.FieldByName("Score").SetFloat(92.5)

  fmt.Println(*student)

}

func adapter(callee interface{}, args ...interface{}) {

  refValue := reflect.ValueOf(callee)
  refType := refValue.Type()

  if refType.Kind() != reflect.Func {
    fmt.Printf("The expected type is a function, but received %v\n", refType)
    return
  }

  // 准备函数调用时的参数
  var params []reflect.Value = nil
  if len(args) > 0 {
    params = make([]reflect.Value, len(args))
    for i := 0; i < len(args); i++ {
      params[i] = reflect.ValueOf(args[i])
    }
  }

  // 使用反射机制调用函数
  retValues := refValue.Call(params)
  if retValues != nil {
    fmt.Println(retValues[0].Int())
  }

}

func reflectStruct(obj interface{}) {

  refValue := reflect.ValueOf(obj)
  refType := refValue.Type()

  if refType.Kind() != reflect.Ptr {
    fmt.Println("The object being reflected should be a pointer")
    return
  }

  if refType.Elem().Kind() != reflect.Struct {
    fmt.Println("The object being reflected should be a struct")
    return
  }

  // 获取struct的所有字段
  fmt.Println("所有的字段如下")
  fmt.Printf("%s\t%s\t%s\n", "字段名", "值", "Tag")
  for i := 0; i < refValue.Elem().NumField(); i++ {
    field := refType.Elem().Field(i)
    fieldValue := refValue.Elem().Field(i)
    tag := field.Tag.Get("json")
    fmt.Printf("%v\t%v\t%v\n", field.Name, fieldValue, tag)
  }

  // 设置struct中field的值
  refValue.Elem().FieldByName("Age").SetInt(22)

  // 列出所有的方法
  fmt.Println("\n所有的方法如下")
  methodCount := refType.NumMethod()
  for i := 0; i < methodCount; i++ {
    method := refType.Method(i)
    fmt.Printf("Method %d: %s\t%v\n", i+1, method.Name, method.Type)
  }

  // 使用反射机制调用方法
  refValue.Method(0).Call([]reflect.Value{reflect.ValueOf(10.0)})

  fmt.Println("\n调用Rename方法，然后显示返回值")
  retValues := refValue.MethodByName("Rename").Call([]reflect.Value{reflect.ValueOf("Timmy")})
  println("the old name is ", retValues[0].Interface().(string))

  fmt.Println("\n调用Print方法")
  refValue.MethodByName("Print").Call(nil)

}
