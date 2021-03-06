对[此文](https://mp.weixin.qq.com/s/Lv2XTD-SPnxT2vnPNeREbg)的验证

环境：

goos: windows

goversion:1.18

goarch: amd64

cpu: AMD Ryzen 5 4600U with Radeon Graphics



## 常用数据结构

### 1. 反射虽好，切莫贪杯

#### 1.1 优先使用 strconv 而不是 fmt

| 代码                                                         | 验证结果                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [strconv_vs_fmt_test.go](datastructure/strconv_vs_fmt_test.go) | [strconv_vs_fmt_test.md](datastructure/strconv_vs_fmt_test.md) |

#### 1.2 少量的重复不比反射差

| 代码                                                         | 验证结果                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [do_not_overuse_reflection_test.go](datastructure/do_not_overuse_reflection_test.go) | [do_not_overuse_reflection_test.md](datastructure/do_not_overuse_reflection_test.md) |

#### 1.3 慎用 binary.Read 和 binary.Write

| 代码                                                         | 验证结果                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [use_binary_read_write_carefully_test.go](datastructure/use_binary_read_write_carefully_test.go) | [use_binary_read_write_carefully_test.md](datastructure/use_binary_read_write_carefully_test.md) |

### 2. 避免重复的字符串到字节切片的转换

| 代码                                                         | 验证结果                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [avoid_unnecessary_initialization_test.go](datastructure/avoid_unnecessary_initialization_test.go) | [avoid_unnecessary_initialization_test.md](datastructure/avoid_unnecessary_initialization_test.md) |

### 3. 指定容器容量

#### 3.1 指定 map 容量提示

| 代码                                                         | 验证结果                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [map_capacity_hint_test.go](datastructure/map_capacity_hint_test.go) | [map_capacity_hint_test.md](datastructure/map_capacity_hint_test.md) |

#### 3.2 指定切片容量

| 代码                                                         | 验证结果                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [slice_capacity_test.go](datastructure/slice_capacity_test.go) | [slice_capacity_test.md](datastructure/slice_capacity_test.md) |

### 4. 字符串拼接方式的选择

#### 4.1 行内拼接字符串推荐使用运算符+

| 代码                                               | 验证结果                                           |
| -------------------------------------------------- | -------------------------------------------------- |
| [join_str_test.go](datastructure/join_str_test.go) | [join_str_test.md](datastructure/join_str_test.md) |

#### 4.2 非行内拼接字符串推荐使用 strings.Builder

| 代码                                                         | 验证结果                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [str_join_builder_buffer_byte_test.go](datastructure/str_join_builder_buffer_byte_test.go) | [str_join_builder_buffer_byte_test.md](datastructure/str_join_builder_buffer_byte_test.md) |

### 5. 遍历 []struct{} 使用下标而不是 range

#### 5.1 []int，index与range基本无差别

| 代码                                                     | 验证结果                                                 |
| -------------------------------------------------------- | -------------------------------------------------------- |
| [iterate_int_test.go](datastructure/iterate_int_test.go) | [iterate_int_test.md](datastructure/iterate_int_test.md) |

#### 5.2 []struct{}，index性能更优

| 代码                                                         | 验证结果                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [iterate_struct_test.go](datastructure/iterate_struct_test.go) | [iterate_struct_test.md](datastructure/iterate_struct_test.md) |

#### 5.3 []*struct，因为有了指针，index与range基本无差别

| 代码                                                         | 验证结果                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [iterate_struct_pointer_test.go](datastructure/iterate_struct_pointer_test.go) | [iterate_struct_pointer_test.md](datastructure/iterate_struct_pointer_test.md) |

## 内存管理

### 1.使用空结构体节省内存

#### 1.1 不占内存空间

| 代码                                                         | 验证结果 |
| ------------------------------------------------------------ | -------- |
| [size_of_empty_struct_test.go](datastructure/size_of_empty_struct_test.go) | 0        |

#### 1.2.1 实现集合（Set）

| 代码                                     | 验证结果                                 |
| ---------------------------------------- | ---------------------------------------- |
| [set_test.go](datastructure/set_test.go) | [set_test.go](datastructure/set_test.go) |

