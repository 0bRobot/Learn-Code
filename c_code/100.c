#define _CRT_SECURE_NO_WARNINGS 1
// C语言入门编程100题

 NO.01
# include <stdio.h> 

// 1 2 3 4 个数字组成多少个互不相同的无重复的三位数字   分别是多少
int main()
{
	int i, j, k;  // 定义三个变量
	int a = 0;  //计数
	printf("\n"); //打印空行
	for (i = 1; i <= 4; i++)  // 以下为三重循环
	{
		for (j = 1; j <= 4; j++)
		{
			for (k = 1; k <= 4; k++)
			{
				if (i != k && i != j && j != k)  // 确保i、j、k三位互不相同
				{
					printf("%d %d %d\n",i ,j,k);
					a++;  // 打印一次 自增1
					
				}
			}
		}
	}
	printf("一共有%d组\n", a);  // 打印总数
	return 0;
}

// NO.02
#include <stdio.h>

 //企业发放的奖金根据利润提成。 从键盘输入当月利润I，求应发放奖金总数？
//利润(I)低于或等于10万元时，奖金可提10% ；
//利润高于10万元，低于20万元时，低于10万元的部分按10% 提成，高于10万元的部分，可提成7.5 % ；
//20万到40万之间时，高于20万元的部分，可提成5 % ；
//40万到60万之间时高于40万元的部分，可提成3 % ；
//60万到100万之间时，高于60万元的部分，可提成1.5 % ；
//高于100万元时，超过100万元的部分按1 % 提成

int main()
{
	double I;  //当月近利润
	double  Bouns, Bouns1, Bouns2, Bouns4, Bouns6, Bouns10;  // 6个不同的类别的提成策略
	Bouns = 1;
	printf("输入你当月的利润：<");
	scanf("%lf", &I);
	Bouns1 = 100000 * 0.1;  //  低于等于 10万 的奖金
	Bouns2 = Bouns1 + 100000 * 0.075;  // 高于10万 低于20万的奖金  =  低于10万的奖金 +  高出10万部分的奖金
	Bouns4 = Bouns2 + 200000 * 0.05;
	Bouns6 = Bouns4 + 200000 * 0.03;
	Bouns10 = Bouns6 + 400000 * 0.015;
	
	if (I <= 100000)
	{
		Bouns = I * 0.1;  //利润低于10万的  奖金=  利润*10%
	}
	else if (I <= 200000)
	{
		Bouns = Bouns1 + (I - 100000) * 0.075;  // 低于等于10万的奖金加上    高出的部分 提成 7.5%  
		
	}
	else if (I <= 400000)
	{
		Bouns = Bouns2 + (I - 200000) * 0.05;
	}
	else if (I < 600000)
	{
		Bouns = Bouns4 + (I - 400000) * 0.03;
	}
	else if (I <= 1000000)
	{
		Bouns = Bouns6 + (I - 600000) * 0.015;
	}
	else if (I > 1000000)
	{
		Bouns = Bouns10 + (I - 1000000) * 0.01;
	}
	printf("你的奖金为：%lf\n",Bouns);
	return 0;
}


//NO.3

#include <stdio.h>

//一个整数，它加上100后是一个完全平方数，再加上168又是一个完全平方数，请问该数是多少？

int main()
{
	int i, j, m, n, x;
	for (i = 1; i < 168 / 2 + 1; i++)
	{
		if (168 % i == 0)
		{
			j = 168 / i;
			if (i > j && (i + j) % 2 == 0 && (i - j) % 2 == 0)
			{
				m = (i + j) / 2;
				n = (i - j) / 2;
				x = n * n - 100;
				printf("%d + 100 = %d * %d\n",x,n,n);
				printf("%d + 268 = %d * %d\n",x,m,m);
			}
		}
	}
	return 0;
}


//NO.4
#include<stdio.h>

//输入某年某月某日，判断这一天是这一年的第几天？
int main()
{
	int day, month, year,leap;
	int sum =0;
	printf("\n请输入年、月、日,格式为：年 月 日（2015 12 10）\n");
	scanf("%d %d %d",&year,&month,&day); 
	switch (month)
	{
	case 1:
		sum = 0;   // 1月不计算
		break;
	case 2:
		sum = 31;  // 2月 要增加 1月的天数
		break;
	case 3:        // 3月  1月和2月的天数之和；
		sum = 59;
		break;
	case 4:
		sum = 90;
		break;
	case 5:
		sum = 120;
		break;
	case 6:
		sum = 151;
		break;
	case 7:
		sum = 181;
		break;
	case 8:
		sum = 212;
		break;
	case 9:
		sum = 243;
		break;
	case 10:
		sum = 273;
		break;
	case 11:
		sum = 304;
		break;
	case 12:
		sum = 334;
		break;
	default:
		printf("date Error\n");
		break;
	}
	sum = sum + day; // 加上某天的天数
	if (year % 400 == 0 || (year % 4 == 0 && year % 100 != 0))  //判断是不是闰年
	{
		leap = 1;
	}
	else
	{
		leap = 0;
	}
	if ((leap == 1) && (month > 2)) //如果是闰年且月份大于2,总天数应该加一天
	{
		sum++;
	}
	printf("这是一年的第%d天\n",sum);

	return 0;
}

//NO.5
#include <stdio.h>

int main()
{
	int a = 0;
	int b = 0;
	int c = 0;
	// a>b>c
	printf("-------------------------------\n");
	printf("请输入三个整数数:\n");
	printf("-------------------------------\n");
	printf("空格分开\n");
	scanf("%d %d %d",&a,&b,&c);

	if ( a > b )  //交换 a和b 的值
	{
		int tmp = 0;
		tmp = a;   
		a = b;    
		b = tmp;   
	}
	if ( a > c )  
	{
		int tmp=0;
		tmp = a;    
		a = c;
		c = tmp;

	}
	if ( b > c ) 
	{
		int tmp = 0;
		tmp = b;
		b = c;
		c = tmp;

	}
	printf("%d  %d  %d", a, b, c);
	return 0;
}


//NO.6
#include<stdio.h>

//用* 号输出字母C的图案
int main()
{
	printf("     *\n");
	printf("  **    **\n");
	printf("**\n");
	printf("*\n");
	printf("**\n");
	printf("  **    **\n");
	printf("     *\n");

	return 0;
}

NO.7
#include<stdio.h>

//输出特殊图案，请在c环境中运行，看一看，Very Beautiful!

int main()
{
    int a, b;
    for (a = 0; a < 5; a++) //列
    {
        for (b = 0; b < 5; b++) //行
        {
            if (a == b || (a + b+1) == 5)
            {
                printf("■");
            }
            else
                printf("□");
            
        }
        printf("\n");
    }
    return 0;
}


//NO.8
#include<stdio.h>
//输出9*9口诀
int main()
{
	int i, j;
	for (i = 1; i <= 9;i++)
	{
		
		for (j = 1; j <= i; j++)
		{
			printf("%d*%d=%-3d",j, i, i*j);
		}
		printf("\n");
	}
	return 0;
}


NO.9
#include <stdio.h>
//要求输出国际象棋棋盘。
int main()
{
	int i,j;
    for (i = 0; i < 8; i++) // 控制列
    {
        for (j = 0; j < 8; j++)  // 控制行
        {
            if ((i + j) % 2 == 0)
                printf("■");
            else 
                printf("□");
        }
        printf("\n");
    }
}




//NO.10
#include<stdio.h>
// 打印楼梯，同时在楼梯上方打印两个笑脸
int main()
{
	int i,j;
	printf("♂♀\n");
	for (i = 1; i <= 11; i++)
	{
		for (j = 0; j < i; j++)
		{
			printf("■■");
		}
		printf("\n");
		
	}
	return 0;
}


// No.11
#include <stdio.h>
int main()
//古典问题（兔子生崽）：有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？（输出前40个月即可）
//兔子的出生规律为 每月   1  ，1  ，3 ，5  ，8   ，13  ，21   即下个月是上两个月之和（从第三个月开始）
//1 月 2月 都有一只   五月的为 三月和四月之和
{
	long f1 = 1, f2 = 1;  // 兔子数量
	int i;  //循环
	int m =0; //计算月份

	printf("输入要计算的月份_:");
	scanf("%d",&m);
	printf("\n");

	if (m % 2 == 1)  // 按照月份计算循环次数
	{
		m = (m+1) / 2;
	}
	else
	{
		m = m / 2;
	}
	for (i = 1;i <= m; i++)
	{
		printf("第%d个月有%d只兔子\n",i*2-1,f1);
		printf("第%d个月有%d只兔子\n", i * 2, f2);
		f1 = f1 + f2; //前两个月加起来赋值给第三个月
		f2 = f1 + f2; //前两个月加起来赋值给第三个月
	}

	return 0;
}


// No.12
#include <stdio.h>
//判断 101 到 200 之间的素数。
// 用一个数分别去除 2 到 sqrt(这个数)，如果能被整除，则表明此数不是素数   只能被1 和 本身整除
int main()
{
	int i, j;
	int count=0;
	for (i = 101; i <= 200; i+=2) //偶数一定不是素数
	{
		for (j = 2; j < i; j++)
			if (i % j == 0) // 判断能否被整除   再跳出循环
				break;
		if (j >= i)   //判断循环是否提前跳出，如果 j < i 说明在 2~j 之间, i 有可整除的数
		{
			count ++;   //计数器
			printf("%d ", i);

			if (count % 5 == 0) //  换行，用 count 计数，每五个数换行
				printf("\n");
		}

	}
	printf("\n");
	printf("100-200 一共有%d 个素数\n",count);
}


// No.13
#include<stdio.h>

int main()
//打印出所有的"水仙花数"
//其各位数字立方和等于该数 本身。例如：153是一个"水仙花数"，因为153=1的三次方＋5的三次方＋3的三次方
{
	int i, x, y, z;
	int count=0;
	for (i = 100;i < 1000; i++) // 生成所有的三位数
	{
		x = i % 10;           //取个数
		y = i / 10 % 10;     //取十数
		z = i / 100 % 10;     //取百数

		if (i == (x * x * x + y * y * y + z * z * z))
		{
			count++;  //计数
			printf("%d ", i);
			
			if (count % 5 == 0) //每五个数换行
				printf("\n");
		}
	}
	printf("一共有%d个水仙花数\n",count);
	return 0;
}



// No.13
#include<stdio.h>

int main()
//打印出所有的"水仙花数"四位数
//其各位数字立方和等于该数 本身
{
	int i, x, y, z,q;
	int count=0;
	for (i = 1000;i < 10000; i++) // 生成所有的三位数
	{
		x = i % 10;           //取个数
		y = i / 10 % 10;     //取十数
		z = i / 100 % 10;     //取百数
		q = i / 1000 % 10;    //取千数
		if (i == (x * x * x * x + y * y * y * y + z * z * z * z + q * q * q* q))
		{
			count++;  //计数
			printf("%d ", i);
			
			if (count % 5 == 0) //每五个数换行
				printf("\n");
		}
	}
	printf("一共有%d个水仙花数\n",count);
	return 0;
}

//No.14
#include<stdio.h>
//将一个正整数分解质因数。例如：输入90, 打印出90 = 2 * 3 * 3 * 5
int main()
{
	int n, j;
	printf("请输入一个整数：");
	scanf("%d",&n);
	printf("%d=",n);

	for (j = 2; j <=n; j++)  // j大于2 小于  n
	{
		while (n % j == 0)     // n 能被j 整除 进入循环
		{
			printf("%d", j);   //打印第一次的商   
			n /= j;  // n = n/j;  n =  n /2  
			if (n !=  1)   // 商不为1   打印 乘号  
				printf("*");
		}
	}
	printf("\n");
	return 0;

}

//No.15
#include <stdio.h>

int main()
//条件运算符的嵌套来完成此题：学习成绩 >= 90分的同学用A表示，60 - 89分之间的用B表示，60分以下的用C表示
//(a > b) ? a : b这是条件运算符的基本例子。
{
	int score;
	char  grade;
	printf("请输入分数：");
	scanf("%d",&score);
	grade = (score >= 90) ? 'A' : ((score >= 60) ? 'B' : 'C');

	printf("%c \n",grade);
	return 0;
}

// No.16
#include <stdio.h>
//输入两个正整数m和n，求其最大公约数和最小公倍数。
int main()
{
	int a, b, r, n;
	printf("请输入两个数字:");
	scanf("%d %d", &a, &b);
	if (a < b)   // 数值交换   a>b;
	{
		int tmp;
		tmp = b;
		b = a;
		a = tmp;
	}
	r = a % b;  //  判断是否可以整除    11
	n = a * b;  // 可以整除 较大值为最大公约数  较小值为公倍数    420
	while (r != 0)  //判断 取模的值是否为0；
	{
		a = b;
		b = r;
		r = a % b;  //
	}
	printf("这两个数的最大公约数是%d，最小公倍数是%d\n", b, n / b);
	return 0;
}

// No.17
#include <stdio.h>

int main()
//输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
{
	char c;  //char 字符型
	int ZM = 0, spaces = 0, Nu = 0, other = 0;
	printf("请输入一写内容\n");
	while ((c = getchar()) != '\n')  //获取输入的值   
	{
		if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'))   //判断字母  getchat  的值 大于 aA   小于 zZ  逻辑与 逻辑或
		{
			ZM++;
		}
		else if (c >= '0' && c<= '9')  //  字符  必须单引号
		{
			Nu++;
		}
		else if (c == ' ')   ////  字符  必须单引号
		{
			spaces++;
		}
		else
			other++;
	}
	printf("字母=%d 数字=%d 空格=%d 其他=%d \n", ZM, Nu, spaces, other);

		
}

//No.18
#include<stdio.h>
int main()
{
    int s = 0, a, n, t;
    printf("请输入 a 和 n：\n");// a  是值  n 是加几次
    scanf("%d%d", &a, &n);
    t = a;
    while (n > 0)
    {
        s += t;
        a = a * 10;  //2+22+222     每加一次 等于*10  再加上原来的a
        t += a;
        n--;
    }
    printf("a+aa+...=%d\n", s);
}

//No.19
#include <stdio.h>
//一个数如果恰好等于它的因子之和，这个数就称为"完数"。例如6=1＋2＋3编程找出1000以内的所有完
#define N 1000

int main()
{
	int i, j, k, n, sum;
	int a[256];
	
	for (i = 2; i <= N; i++)  //遍历 2 -1000 的数
	{
		sum = a[0] = 1;
		k = 0;
		for (j = 2 ; j <= (i / 2); j++)
		{
			if (i % j == 0)
			{
				sum += j;
				a[++k] = j;
			}
		}
		if (i == sum)
		{
			printf("%d=%d", i, a[0]);
			for (n = 1; n <= k; n++)
				printf("+%d", a[n]);
			printf("\n");
		}
	}
	return 0;
}

//No.20

#include <stdio.h>

int main()
{
	float h, s;
	h = s = 100;
	h /= 2;  //第一次反弹高度
	for (int i = 2; i <= 2; i++)   //控制反弹次数
	{
		s = s + 2 * h; 
		h /=2;  //高的每次减少一半
	}
	printf("10次落地经过%f米，第10次反弹高度%f", s, h);
	return 0;
}

//NO.21
#include <stdio.h>

int main()
// 猴子第一天摘下若干个桃子，当即吃了一半，还不瘾，又多吃了一个
//第二天早上又将剩下的桃子吃掉一半，又多吃了一个。以后每天早上都吃了前一天剩下
//的一半零一个 到第10天早上想再吃时，见只剩下一个桃子了。求第一天共摘了多少
//x1为前一天桃子数，设x2为第二天桃子数， 则：x2=x1/2-1, x1=(x2+1)*2 
{
	int x = 1, day = 9;
	while (day > 0)
	{
		x = (x +1)*2; //第二天加1 的2倍是前一天的数量
		day--;

	}
	printf("%d", x);
	return 0;
}

int main()
{
	int x = 1534, day = 9, n;
	while (day<10&&day!=0)
	{
		x = x / 2 - 1;
		day--;

	}
	return 0;
}

//No.22
#include<stdio.h>
//
int main()
// 个乒乓球队进行比赛，各出三人。甲队为a,b,c三人，乙队为x,y,z三人。已抽签决定比赛名单。有人向队员打听比赛的名单。a说他不和x比，c说他不和x,z比，
{
	char  a, b, c;
	for (a = 'x'; a <= 'z'; a++)
	{
		for (b = 'x'; b < 'z'; b++)
		{
			if (a != b)
			{
				for (c = 'x'; c < 'z'; c++)
				{
					if (a != c && b != c)
					{
						if (a != 'x' && c != 'x' && c != 'z')  // a说他不和x比，c说他不和x,z比
							printf("a->%c  b->%c  c->%c",a,b,c);
					}
				}
			}

		}
	}

	return 0;
}

No.23
#include <stdio.h>
//打印菱形
int main()
{
	int i, j, k;
	for (i = 0; i <= 3; i++)  //循环控制打印  上面四行   控制行
	{
		for (j = 0; j <= 2 - i; j++)     // 控制输入空格
		{
			printf(" ");
		}
		for (k = 0; k <= 2 * i; k++)   //控制输出*
		{
			printf("*");
		}
		printf("\n");    
	}
	for (i = 0; i <= 2; i++)  //循环控制打印  下面四行 
	{
		for (j = 0; j <= i; j++)
		{
			printf(" ");
		}
		for (k = 0; k <= 4 - 2 * i; k++)
		{
			printf("*");
		}
		printf("\n");
	}
	return 0;
}

//No.24
#include <stdio.h>
//有一分数序列：2/1，3/2，5/3，8/5，13/8，21/13...求出这个数列的
//前20项之和
int main()
{
	int  t;
	float sum = 0;
	float a = 2, b = 1;
	for (int i=1;i<=20;i++)
	{
		sum = sum + a / b;  // a/b 的值赋值给sum   
		t = a;   // a 的赋值给 临时变量t
		a = a + b;  // a+b 的值 赋值给a
		b = t;  // 临时变量t 赋值给b  
	}
	printf("%9.6f\n",sum);
	return 0;
}

//No.25
#include <stdio.h>
// 求1+2!+3!+...+20!的和。  
// 及 1+ 2*1+3*2*1+4*3*2*1 ...
int main()
{
	long double sum=0,mix=1;
	for (int i = 1; i <= 4; i++)
	{
		mix = mix * i;
		sum = sum + mix;
	}
	
	printf("%Lf",sum);  //%Lf 是输出 long double 型变量   %f 是输出 double 型变量 %f 是输出 float 型变量
	return 0;
}

//No.26
#include <stdio.h>
// 利用递归方法求5!。  
int fact(int j)
{
	int sum;
	if (j == 0)
	{
		sum = 1;
	}
	else
	{
		sum = j * fact(j - 1);
	}
	return sum;
}

int main()
{
	int fact(int);
	for(int i = 0; i <= 5; i++)
	{
		printf("%d!=%d\n",i,fact(i));
	}
	return 0;
}

//No.27
#include <stdio.h>
//利用递归函数调用方式，将所输入的5个字符，以相反顺序打印出来
void Prinf(n)
int n;
{
	char next;
	if (n <= 1)
	{
		next = getchar();
		printf("输出相反顺序\40:\40");
		putchar(next);
	}
	else
	{
		next = getchar();
		Prinf(n - 1);
		putchar(next);
	}
}

int main()
{
	int i = 5;
	void Prinf(int n);
	printf("请输入5个字符\40:\40");
	Prinf(i);
	printf("/n");
	return 0;
}

// No.28
//有5个人坐在一起，问第五个人多少岁？他说比第4个人大2岁。问第4个人岁数，他说比第3个人大2岁。问第三个人，又说比第2人大两岁。问第2个人，说比第一个人大两岁。最后问第一个人，他说是10岁。请问第五个人多大
#include <stdio.h>
int age(n)
int n;
{
	int c;
	if (n == 1) c = 10;
	else
	{
		c = age(n - 1) + 2;
		}
	return(c);
}

int main()
{
	printf("%d\n", age(5));
	return 0;
}

// No.29 
//给一个不多于5位的正整数，要求：一、求它是几位数，二、逆序打印出各位数字。
#include<stdio.h>
int main()
{
	long a, b, c, d, e, x;
	printf("请输入5个字符：");
	scanf("%ld", &x);
	a = x / 10000; //求的万位
	b = x % 10000 / 1000;  //求的千位
	c = x % 1000 / 100;  //求的百位
	d = x % 100 / 10; //求的十位
	e = x % 10; //求的个位

	if (a != 0) {
		printf("为 5 位数,逆序为： %ld %ld %ld %ld %ld\n", e, d, c, b, a);
	}
	else if (b != 0) {
		printf("为 4 位数,逆序为： %ld %ld %ld %ld\n", e, d, c, b);
	}
	else if (c != 0) {
		printf("为 3 位数,逆序为：%ld %ld %ld\n", e, d, c);
	}
	else if (d != 0) {
		printf("为 2 位数,逆序为： %ld %ld\n", e, d);
	}
	else if (e != 0) {
		printf("为 1 位数,逆序为：%ld\n", e);
	}
	return 0;
}


// No.30
#include <stdio.h>
// 列出所有五位数的回文数
// 一个5位数，判断它是不是回文数。即12321是回文数，个位与万位相同，十位与千位相同
int main()
{
	int count = 0;
	for (int i = 10000; i <= 99999; i++)
	{
		int w=0, q=0, s=0, g=0;
		w = i / 10000;   // 求万位的值
		q = i % 10000 / 1000;  // 求千位的值
		s = i % 100 / 10;   // 求十位的值
		g = i % 10;     // 求个位的值
		
		if(w==g&&q==s)  //判断 万位和个位   千位和十位是否相等
		{
			count++;
			printf("%d  ",i);

			if (count % 9 == 0) //  换行，用 count 计数，每五个数换行
				printf("\n");

		}
	}
	printf("一共有%d个回文数",count);
	return 0;
}

//NO.31

#include<stdio.h>
//请输入星期几的第一个字母来判断一下是星期几，如果第一个字母一样，则继续判断第二个字母


int main()
{
	char i, j;
	printf("请输入星期的第一个字母：\n");
	scanf("%c",&i);
	getchar();  //getchar()读取缓存
	switch (i)
	{
	case 'm':
		printf("monday  1 \n");
		break;
	case 'w':
		printf("wednesday  3\n");
		break;

	case 'f':
		printf("friday 5  \n");
		break;
	
	case 't':
		printf("请输入下一个字母：\n");
		scanf("%c", &j);
		if (j == 'h')
		{
			printf("thursday   4\n");
			break;
		}
		if (j == 'h')
		{
			printf("thesday  2\n");
			break;
		}
	case 's':
		printf("请输入下一个字母：");
		scanf("%c",&j);
		if (j == 'a')
		{
			printf("saturday 6\n");
			break;
		}
		if (j == 'u')
		{
			printf("sunday  7\n");
			break;
		}
	default:
		printf("输入错误\n");
		break;
	}

	return 0;

}

#include <stdio.h>
#include <math.h>


int main()
{
	float a, b, c,x;
	float K, S, Q, Y,Z;

	printf("请输入三个值：<");
	scanf("%f %f %f",&a,&b,&c);
	getchar();
	if (a == 0 && b == 0)
	{
		printf("Data error\n");
	
	}
	if (a == 0 && b != 0)
	{
		x = -c / b;
		printf("%f", x);
	}
	K = pow(b, 2) - 4.0 * a * c;
	if (K < 0)
	{
		printf("无解");
	}
	else if( K==0 )
	{
		S = -b / (2 * a);
		printf("%f", S);
	}
	else
	{
		Q = sqrt(K);
		Y = (-b + Q) / (2.0 * a);
		Z = (-b - Q) / (2.0 * a);
		printf("%f%f", Y, Z);
	}


	return 0;
}

//NO.32
// 删除一个字符串中的指定字母，如：字符串 "aca"，删除其中的 a 字母。

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* DelChar(char* sdel, char* sCH)
{
	int AA[256];
	if ( sdel == NULL )  //判断删除字母是否为空
		return sCH;      // 为空返回 sCH

	for (int i = 0; i < 256; i++)
		AA[i] = 0;
	for (int i = 0; i < strlen(sdel); i++) // 求删除字符串长度 是否为1
		AA[sdel[i]] = 1;
	
	int bb=0;
	for (int i = 0; i < strlen(sCH); i++) //先遍历目标字符串长度
	{
		if (!AA[sCH[i]])
			sCH[bb++] = sCH[i];
	}
	sCH[bb] = '\0';
	return sCH;
}

int main()
{
	char del[2] = "a";   //要删除的字母
	char CH[7] = "abcdard"; // 目标字符串
	printf("%s\n", DelChar(del, CH));
	return 0;
}

//No33
#include<stdio.h>
#include<math.h>
#define MAX 1000

//判断一个数字是否为质数 /素数

int prime[MAX];

int isPrimeNaive(int n)
{
    if (n <= 1)
        return 0;
    for (int i = 2; i < n; i++)
        if (n % i == 0)
            return 0;
    return 1;
}

int isPrime(int n)
{
    if (n <= 1)
        return 0;
    if (n == 2)
        return 1;
    if (n % 2 == 0)
        return 0;
    int limit = (int)sqrt((double)n);
    for (int i = 3; i <= limit; i = i + 2)
    {
        if (n % i == 0)
            return 0;
    }
    return 1;
}

void sieve()
{
    prime[0] = 0;
    prime[1] = 0;
    for (int i = 2; i < MAX; i++)
        prime[i] = 1;
    int limit = (int)sqrt((double)MAX);
    for (int i = 2; i <= limit; i++)
    {
        if (prime[i])
            for (int j = i * i; j <= MAX; j += i)
                prime[j] = 0;
    }
}

int isPrimeSieve(int n)
{
    if (prime[n])
        return 1;
    else
        return 0;
}

int main()
{
    sieve();
    printf("N=%d %d\n", 1, isPrime(1));
    printf("N=%d %d\n", 2, isPrime(2));
    printf("N=%d %d\n", 3, isPrime(3));
    printf("N=%d %d\n", 4, isPrime(4));
    printf("N=%d %d\n", 7, isPrime(7));
    printf("N=%d %d\n", 9, isPrime(9));
    printf("N=%d %d\n", 13, isPrime(13));
    printf("N=%d %d\n", 17, isPrime(17));
    printf("N=%d %d\n", 100, isPrime(100));
    printf("N=%d %d\n", 23, isPrime(23));
    printf("N=%d %d\n", 1, isPrime(1));
    return 0;
}

// No 34
#include <stdio.h>
// 函数的调用
void hello_p()  //被调用的函数
{
	printf("hello world！\n");
}

void three_hellos()
{
	int a;
	for (a = 1; a <= 5; a++)
		hello_p();//调用函数
}

int main()
{
	three_hellos(); //调用函数
}

#include <stdio.h>
// 字符串反转

void reverse(char* s)
{
	int len = 0;
	char* p = s;
	while (*p != 0)
	{
		//printf("%s", p);
		len++;         //求字符串长度
		p++; //指针自增   
	}

	int i = 0;
	char c;
	while (i <= len / 2 - 1)
	{  //交换字符串
		c = *(s + 1);
		*(s + i) = *(s + len - 1 - i);
		*(s + len - 1 - i) = c;
		i++;

	}

}
int main()
{
	char s[] = "nihaoa!";
	printf("%s \n", s);
	reverse(s);// 反转字符串
	printf("%s\n", s);
	return 0;
}

#include <stdio.h>
#include <math.h>
//求100以内的素数
int main()
{
    int i, j, k, n = 0;
    for (i = 2; i <= 100; i++)
    {
        k = (int)sqrt(i);
        for (j = 2; j <= k; j++)
            if (i % j == 0) break;
        if (j > k)
        {
            printf("%d ", i);
            n++;
            if (n % 5 == 0)
                printf("\n");
        }
    }
    return 0;
}

#include<stdio.h>
#define N 3
int main()
{
    int i, j, a[N][N], sum = 0,sum1=0;  // 3 行 3 列
    printf("请输入矩阵(3*3)：\n");
    
    for (i = 0; i < N; i++)
        for (j = 0; j < N; j++)
            scanf("%d", &a[i][j]);
    
    for (i = 0; i < N; i++)
    
        sum += a[i][i]; //  a[0][0] + a [1][1] + a [2][2]


    printf("对角线之和为：%d\n", sum);
    printf("对角线之和为：%d\n", sum1 = a[0][2] + a[1][1] + a[2][0]);

    return 0;
}

#include <stdio.h>
//No 39
//有一个已经排好序的数组。现输入一个数，要求按原来的规律将它插入数组中
int main()
{
    int a[12]={1,4,6,9,13,16,19,28,40,65,100};
    int temp1,temp2,number,end,i,j;
    printf("原始数组是:\n");
    for(i=0;i<11;i++)
        printf("%4d",a[i]);
    printf("\n插入一个新的数字: ");
    scanf("%d",&number);
    end=a[10];
    if(number>end)
        a[10]=number;
    else
    {
        for(i=0;i<11;i++)
        {
            if(a[i]>number)
            {
                temp1=a[i];
                a[i]=number;
                for(j=i+1;j<12;j++)
                {
                    temp2=a[j];
                    a[j]=temp1;
                    temp1=temp2;
                }
                break;
            }
        }
    }
    for(i=0;i<12;i++)
        printf("%4d",a[i]);
    printf("\n");
    return 0;
}

// No_40
//将一个数组逆序输出
#include<stdio.h>
#define N 11
int main()
{
    int a[N] = { 0,1,2,3,4,5,6,7,8,9,10};
    int i, t;
    
    printf("原始数组是:\n");
    for (i = 0; i < N; i++)
        printf("%d ", a[i]);


    for (i = 0; i < N / 2; i++)  // 用第一个与最后一个交换
    {
        t = a[i];
        a[i] = a[N - 1 - i];
        a[N - 1 - i] = t;
    }
   
    printf("\n排序后的数组:\n");
    for (i = 0; i < N; i++)
        printf("%d ", a[i]);
    printf("\n");
    return 0;
}
