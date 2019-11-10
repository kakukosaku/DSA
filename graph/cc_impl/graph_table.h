//
// Created by kaku on 2019/11/5.
//

#ifndef DSA_GRAPH_TABLE_H
#define DSA_GRAPH_TABLE_H

#define MaxVertexNum 100

typedef int VertexType

typedef struct ArcNode {
    // 顶点位置
    int adjvex;
    struct ArcNode *next;
} ArcNode;

typedef struct VNode {
    VertexType data;
    ArcNode *first;
}VNode, AdjList[MaxVertexNum;
// adjacent 邻近的

typedef struct {
    AdjList vertices;
    int vexnum, arcnum;
}ALGraph;

#endif //DSA_GRAPH_TABLE_H
