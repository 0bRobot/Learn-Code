//#define _CRT_SECURE_NO_WARNINGS
#include <Windows.h> //Windows  API
#include "resource.h"
//  main()  控制台 入口函数
//WINAPI  相当于 man函数的入口
// HINSTANCE  应用程序实例句柄类型  ID

//窗口处理函数 声明
LRESULT WINAPI WindowProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam);

//绘图函数 声明
void On_Paint(HDC U_Hdc);
//绘制方块
void DrawBlock(HDC U_Hdc);

//游戏地图
//二维数据实现俄罗斯游戏
// 0 表示空白    1  表示 下落的方块   2 已经落下的方块
int g_background[20][10] =  // 20.行 10 列
{
	{0,0,0,0,1,0,0,0,0,0},
	{0,0,0,1,1,1,0,0,0,0}
};

int WINAPI WinMain (HINSTANCE hInstance,HINSTANCE hPreInstance,LPSTR lpCmdLine,int nCmdShow ) //包含4个参数    入栈方式
{
	//弹出一个消息提示框
	//MessageBox(NULL, L"这是我的第一个窗口程序", L"提示", MB_YESNO); //MB_YESNO 选择是否
	//设计窗口类
	
	// 1. 设计窗口类
	TCHAR szAppClassName[] = TEXT("MRGUI");  // TCHAR  宽字符  TEXT  自动适应字符集
	WNDCLASS wc;  //
	wc.cbClsExtra = 0; //窗口类的额外空间大小
	wc.cbWndExtra = 0; //窗口扩展空间大小
	wc.hbrBackground = (HBRUSH)GetStockObject(WHITE_BRUSH); //加载白色背景画刷
	wc.hCursor = LoadCursor(NULL, IDC_ARROW); //加载光标
	wc.hIcon = LoadIcon(hInstance,MAKEINTRESOURCE(IDI_ICON1)); //加载图标
	wc.hInstance = hInstance;   //当前应用程序实例句柄
	wc.lpfnWndProc = WindowProc; //窗口处理函数
	wc.lpszClassName = szAppClassName; //窗口类型名
	wc.lpszMenuName = NULL;   //菜单名 
	wc.style = CS_HREDRAW | CS_VREDRAW;   // 窗口类的风格

	
	//2 注册窗口类
	RegisterClass(&wc);
	
	//3 创建窗口
	HWND hWnd = CreateWindow(
		szAppClassName,       //窗口类型名
		TEXT("俄罗斯方块"),
		WS_OVERLAPPEDWINDOW, //窗口类风格
		300,200,          //窗口左上角坐标
		600,820,           // 窗口宽和高
		NULL,	//父窗口句柄
		NULL,			//菜单句柄
		hInstance,       //应用程序实例句柄
		NULL			//参数
		);
	// 4. 显示窗口
	ShowWindow(hWnd, SW_SHOW);
	
	//5. 更新窗口
	UpdateWindow(hWnd);


	// 6.消息循坏  // while  sellp 导致程序不往下进行
	MSG U_msg;
	while (GetMessage(&U_msg, NULL, 0, 0)) //发送WM_QUIt 才能退出
		//接受返回  接受循坏
	{
		//将虚拟按键消息装换为字符消息
		TranslateMessage(&U_msg);
		//将消息分发给窗口处理函数处理
		DispatchMessage(&U_msg);
	}
	

	return 0;
	
}

//窗口处理函数
//windows 通过消息机制驱动   不像控制台一样执行
LRESULT WINAPI WindowProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam)
{
	PAINTSTRUCT U_PS;
	HDC U_Hdc;
	switch (uMsg)
	{
	case WM_PAINT: //窗口绘图消息   可以在窗即上绘图
	U_Hdc =	BeginPaint(hWnd,&U_PS); //开始画图  在那个窗口上。数据保存到 U_PS结构体中
	On_Paint(U_Hdc); //绘图函数
		
		//绘制矩形
		//Rectangle(U_Hdc,10,10,200,100);  //客户区坐标系  向下 y
		//绘制圆
		//Ellipse(U_Hdc,80,80,200,200); //内切圆
		
		EndPaint(hWnd,&U_PS);//结束绘图
		break;
	
	case WM_CLOSE: //窗口关闭消息  右上角关闭 是VM_CLOSE
		break;
	case WM_DESTROY:   // 窗口销毁消息
		PostQuitMessage(0);// 返回0  为假  
		break;

	}
	return DefWindowProc(hWnd, uMsg, wParam, lParam);
}

//绘图实现
void On_Paint(HDC U_Hdc)  //U_Hdc  界面显示
{
	//双缓冲  
	//方块移动 重新绘制  会闪烁
	HDC HMenDC = CreateCompatibleDC(U_Hdc);
	HBITMAP hBackBMP =  CreateCompatibleBitmap(U_Hdc, 600, 820);  //内存 备件 全黑
	SelectObject(HMenDC, hBackBMP);
	DrawBlock(HMenDC);

	//从内存复制到界面
	BitBlt(U_Hdc, 0, 0, 600, 820, HMenDC, 0, 0, SRCCOPY);

}

//绘制方块
void DrawBlock(HDC U_Hdc)
{
	//绘制一个矩形  游戏区
	Rectangle(U_Hdc,0,0,380,820);
	
	// 绘制每一个方块
	for (int i = 0; i < 20; i++)
	{
		for (int j = 0; j < 10; j++)
		{
			if (g_background[i][j] == 1 || g_background[i][j] == 2)
			{
				//绘制矩形
				Rectangle(U_Hdc,j*30+1,i*30+1,j*30+30-1,i*30+30-1);
			}
		}
	}
	
}