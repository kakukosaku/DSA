cmake_minimum_required(VERSION 3.15)
project(DSA C)

set(CMAKE_C_STANDARD 11)

include_directories(link_list/c_implement)
include_directories(seq_list/c_implement)

add_executable(test_array sort/c_implement/test_array.c sort/c_implement/array.c)
add_executable(
        test_sort
        sort/c_implement/test_sort.c
        sort/c_implement/array.c
        sort/c_implement/bubbl_sort.c
        sort/c_implement/select_sort.c
)