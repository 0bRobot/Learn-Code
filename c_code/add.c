#define  _CRT_SECURE_NO_WARNINGS 1
#include <stdio.h>

void print()
{
	printf("####################################\n");
	printf("------------计算两个数的和----------\n");
	printf("######==========================####\n");
	printf("\n");
	printf("请输入两个数 [空格分开]：");
}

void ADD(long x,long y)
{
	printf("两个数的和为：%2d\n",x+y);
}

int main()
{
	print();
	long muna = 0;
	long munb = 0;
	while (1)
	{
		if ((scanf("%d %d", &muna, &munb)) == 2)
		{
			ADD(muna,munb);
			break;
		}
		else
		{
			printf("!!!!!!!!!!!!!!!!!\n");
			printf("=====是不是瞎重新输入>:\n");
			int c;
			while ((c = getchar()) != '\n' && c != EOF);
		}

	}

	return 0;
}