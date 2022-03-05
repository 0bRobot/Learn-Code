#define _CRT_SECURE_NO_WARNINGS 1
//// C语言入门编程100题

 //NO.01
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

//NO.7
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


//NO.9
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
