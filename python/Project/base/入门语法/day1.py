"""
变量
变量的定义
变量名称  =  变量的值
"""
money = 1000.00
print("我的余额还有:", money, "元")
# 消费了200
money = money - 200
print("消费后剩余: ", money, "元")

# type() 函数 输出数据类型

print(type("名字"))
print(type(13.13))

name = "憨憨"
age = 23
job = "老师"
print("money的类型为:", type(money), "\n"
      "name的类型为：", type(name),"\n"
      "name的类型为：", type(age))

# python 中变量没有类型       变量中存储的数据有类型  如 int float
# 字符串变量   表示变量存储的数据类型是字符串

# python  int  float  string  之间的类型转换

# 对应的转换函数
# int(x)   将x 转换为一个int类型
# float(x)  将x 转换为一个float类型
# str(x)  将x 转换为一个str类型   任何类型可转为为字符串

a = "123"  # 定义一个 str 变量
print(type(a))
print(type(int(a)))
# <class 'int'>  转换为 int 类型
b = 12
print(type(b))
print(type(float(b)), b)   # 第二个b 为什么是整型
b1 = float(b)
print(b1)


c = 2323424
print(type(c))
print(type(str(c)))

# 标识符
"""
变量的名称 方法的名称  类的名称  统称为标识符

python 中标识符命名规则  3类
1 内容限制
   只允许 出现 英文 中文  数字 _     
   不推荐使用中文
   不可以以数字起始
   
2 大小写敏感

3  不可使用关键字
    关键字 大小写敏感   其中 False  Ture   None  为 首字母大写
    
------------------------------------------------------------------------------------------

变量命名规范 
    见名知意
    下划线命名法
    英文字母全部小写    
"""




