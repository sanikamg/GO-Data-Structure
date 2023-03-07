#include<stdio.h>
struct Emp{
    int emp_no;
    char ename[20];
    float esal;

};
void main(){
    char*  cp;
    int* ip;
    struct Emp* sp;
    printf("Integer pointer size : %d",sizeof(sp));
    

   return 0; 
}