cmake_minimum_required(VERSION 3.15)
project(DSA)

set(CMAKE_CXX_STANDARD 14)

include_directories(linked_list/c_implement)
include_directories(list/cc_impl)
include_directories(sort/c_impl)

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
        test_list
        list/cc_impl/test_list.cc
        list/cc_impl/list.cc
)
add_executable(
        test_array
        sort/c_impl/test_array.c
        sort/c_impl/array.c
)
add_executable(
        test_linked_list
        linked_list/cc_impl/test_linked_list.cc
        linked_list/cc_impl/linked_list.cc
)
add_executable(
        test_binary_tree
        tree/cc_impl/test_binary_tree.cc
        tree/cc_impl/binary_tree.cc
)
add_executable(two_sum leetcode/1_two_sum/two_sum.c)
add_executable(add_two_numbers leetcode/2_add_two_numbers/add_two_numbers.cc)
add_executable(
        test_binary_search_tree
        tree/cc_impl/test_binary_search_tree.cc
        tree/cc_impl/binary_search_tree.cc
)