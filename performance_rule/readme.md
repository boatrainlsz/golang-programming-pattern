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
