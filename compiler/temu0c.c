#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>
#include <stdbool.h>


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


    bool parsing = false;
    char* start; int len;

    while(*file_data) {
        // Skip comment lines
        if(*file_data == ';')
            while(*file_data && *file_data != '\n')
                file_data++;

        if(!isspace(*file_data)) {
            if(!parsing) {
                parsing = true;
                start = file_data;
                len = 0;
            }
            len++;
        } else {
            if(parsing) {
                parsing = false;
                if(len != 3) {
                    printf("label: ");
                }
                while(len-- > 0) {
                    printf("%c", *start++);
                }
                printf("\n");
            }
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