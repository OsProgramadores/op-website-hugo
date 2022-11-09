#include <iostream>


int main()
{
	
	for(int i=0; i <= 1000; i++)
	{
		int count = 0;
		
		for( int j=1; j<=i; j++)
	    {
			if(i % j == 0)
			count++;
		}
		
		if(count == 2)
		std::cout << "Numero " << i << " eh primo" << '\n';	 
	}
	
		
}
