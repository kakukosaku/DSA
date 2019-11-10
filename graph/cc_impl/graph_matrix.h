//
// Created by kaku on 2019/11/5.
//

#ifndef DSA_GRAPH_H
#define DSA_GRAPH_H

#define MaxVertexNum 100

typedef char VertexType;

typedef int EdgeType;

typedef struct {
    VertexType Vex[MaxVertexNum];
    EdgeType Edge[MaxVertexNum][MaxVertexNum];
    int vexnum, arcnum;
} MGraph;
// Matrix 矩阵 :)

#endif //DSA_GRAPH_H
