#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>


static uint16_t bin_data[0xFFFF] = {0,};
static uint16_t bin_data_ptr = 0;

int main() {
    // Open file for reading
    FILE* file = fopen("source.asm", "rb");
    if(file == NULL)
        perror ("Error opening source file");

    // Get file size
    fseek(file, 0, SEEK_END);
    long file_size = ftell(file);

    // Read file contents into file_data
    rewind(file);
    char* file_data = malloc(file_size + 1);
    fread(file_data, file_size, 1, file);

    // Close file
    fclose(file);

    while(*file_data) {
        if(*file_data == ';')
            while(*file_data && *file_data != '\n')
                file_data++;

        int _parsing = 0;
        while(*file_data && !isspace(*file_data) && *file_data != ';') {
            printf("%c", *file_data);
            file_data++;
            _parsing = 1;
        }
        if(_parsing) {
            printf("\n");
        }

        file_data++;
    }


    file = fopen("out.bin", "wb+");
    if(file == NULL)
        perror ("Error creating output file");
    fputs((char*)bin_data, file);
    fclose(file);

    return 0;
}