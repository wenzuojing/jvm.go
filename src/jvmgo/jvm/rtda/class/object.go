package class

import (
    . "jvmgo/any"
)

// object
type Obj struct {
    class   *Class
    fields  Any // []Any for Object, []int32 for int[] ...
    extra   Any // todo
}

func (self *Obj) String() string {
    return "{Obj class:" + self.class.String() + "}"
}

// getters
func (self *Obj) Class() (*Class) {
    return self.class
}
func (self *Obj) Fields() (Any) {
    return self.fields
}
func (self *Obj) Extra() (Any) {
    return self.extra
}

func (self *Obj) IsInstanceOf(class *Class) (bool) {
    if class.IsInterface() {
        for k := self.class; k != nil; k = k.superClass {
            for _, i := range k.interfaces {
                if _interfaceXextendsY(i, class) {
                    return true
                }
            }
        }
    } else {
        for k := self.class; k != nil; k = k.superClass {
            if k == class {
                return true
            }
        }
    }
    return false
}
func _interfaceXextendsY(x, y *Class) (bool) {
    if x == y {
        return true
    }
    for _, superInterface := range x.interfaces {
        if _interfaceXextendsY(superInterface, y) {
            return true
        }
    }
    return false
}

// todo
func (self *Obj) zeroFields() {
    fields := self.fields.([]Any)
    for class := self.class; class != nil; class = class.superClass {
        for _, f := range class.fields {
            if !f.IsStatic() {
                fields[f.slot] = f.defaultValue()
            }
        }
    }
}

// reflection
func (self *Obj) GetFieldValue(fieldName, fieldDescriptor string) Any {
    field := self.class.GetField(fieldName, fieldDescriptor)
    return field.GetValue(self)
}
func (self *Obj) SetFieldValue(fieldName, fieldDescriptor string, value Any) {
    field := self.class.GetField(fieldName, fieldDescriptor)
    field.PutValue(self, value)
}