对[此文](https://mp.weixin.qq.com/s/Lv2XTD-SPnxT2vnPNeREbg)的验证

环境：

goos: windows

goversion:1.18

goarch: amd64

cpu: AMD Ryzen 5 4600U with Radeon Graphics

目录：

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

