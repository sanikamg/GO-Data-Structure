#include<stdio.h>

int main(){
    int arr[5],i;

    printf("Enter your 5 Array values : \n");
    for(i=0;i<5;i++){
        scanf("%d",&arr[i]);
    }

    printf("your array values is : \n");
     for(i=0;i<5;i++){
        printf("%d ",arr[i]);
    }
    return 0;
}