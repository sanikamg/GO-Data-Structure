#include<stdio.h>

int main(){
    int arr[5],i;
    int large;
    large=arr[0];

    printf("Enter your 5 Array values : \n");
    for(i=0;i<5;i++){
        scanf("%d",&arr[i]);
    }

    printf("largest element  is : \n");
     for(i=1;i<5;i++){
        if(arr[i]>large) {
            large=arr[i];

        }
        
    }
    printf("%d ",large);
    return 0;
}