在Go语言的反射中，`reflect.Type`和 `reflect.Value`提供了丰富的方法来操作类型和值。下面这个表格汇总了它们的主要方法，并标注了各自的侧重点和共通点。

| **方法类别** | **方法名 (reflect.Type)** | **方法名 (reflect.Value)** | **作用说明 / 区别** |
| --- | --- | --- | --- |
| **获取基础信息** | `Name()`, `String()` | `String()`, `Kind()` | **Type** 侧重类型标识（如类型名）。**Value** 的 `String()`通常返回值的概括信息而非具体值。 |
| **获取种类 (Kind)** | `Kind()` | `Kind()` | **作用相同**。返回基础类型的分类（如 `Int`, `Slice`, `Struct`），这是两者最核心的共用方法之一。 |
| **结构体操作** | `NumField()`, `Field(i)`, `FieldByName()` | `NumField()`, `Field(i)`, `FieldByName()` | **作用相似，返回结果类型不同**。**Type** 的方法返回 `StructField`（包含字段名、类型、标签等元数据）。**Value** 的方法返回 `reflect.Value`（代表该字段的实际值）。 |
| **方法操作** | `NumMethod()`, `Method(i)`, `MethodByName()` | `NumMethod()`, `Method(i)`, `MethodByName()` | **作用相似，返回结果类型不同**。**Type** 返回 `Method`（包含方法类型信息）。**Value** 返回 `reflect.Value`（可用来动态调用方法，如使用 `Call()`)。 |
| **容器类型操作** | `Key()`, `Elem()` | `MapKeys()`, `MapIndex()`, `Len()`, `Index(i)`, `Elem()` | **Type** 的方法（如 `Elem()`）用于查询容器（如切片、映射、指针）内部元素的**类型信息**。**Value** 的方法用于操作具体的**值**，如获取映射的所有键值对、获取切片某个位置的元素值等。 |
| **值操作 (Value独有)** | - | `Int()`, `String()`, `Bool()`, `Interface()`等 | **Value 特有**。用于从反射对象中提取具体的值。使用时必须确保 `Kind()`与该方法匹配。`Interface()`方法可将 `reflect.Value`转换回 `interface{}`，是反射的“逆操作”。 |
| **设值操作 (Value独有)** | - | `SetInt()`, `SetString()`, `Set()`等 | **Value 特有**。用于修改值。前提是该 `Value`是**可寻址的**（通常需要传入变量的指针，并使用 `Elem()`获取指向的值）。 |
| **类型检查 (Type独有)** | `Implements(u Type)`, `AssignableTo(u Type)`等 | - | **Type 特有**。用于高级类型检查，如判断一个类型是否实现了某个接口。 |