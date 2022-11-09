#include <iostream>
#include <string>


int main()
{
	std::string converter, numero_inverso;
	
	for(int i = 0; i <= 64;  i++)
	{
		if(i > 9)		
		converter = std::to_string(i);
		
		if(!converter.empty())
		{
			numero_inverso = converter[1];
			numero_inverso += converter[0];
			
			if(converter == numero_inverso)
			{
				std::cout << "Os numeros palindromicos sao: " << converter << '\n';
			}
	    }		
	}
}
