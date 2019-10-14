cmake_minimum_required(VERSION 3.15)
project(DSA)

set(CMAKE_CXX_STANDARD 14)

include_directories(linked_list/c_implement)
include_directories(seq_list/c_implement)
include_directories(sort/cpp_impl)


add_executable(
        test_array
        sort/c_impl/array.c
        sort/c_impl/test_array.c
)
add_executable(
        test_sort
        sort/c_impl/array.c
        sort/c_impl/test_sort.c
        sort/c_impl/bubbl_sort.c
        sort/c_impl/quick_sort.c
        sort/c_impl/select_sort.c
        sort/c_impl/heap_sort.c
        sort/c_impl/insert_sort.c
        sort/c_impl/shell_sort.c
        sort/c_impl/merge_sort.c
)
add_executable(
        test_array_cpp
        sort/cpp_impl/test_array.cpp
        sort/cpp_impl/array.cpp
)