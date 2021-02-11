package BinaryTree

import(
  "fmt"
)

type Node struct{
  value int
  left  *Node
  right *Node
}

type BinaryTree struct{
  root *Node
}


func (binaryTree *BinaryTree) Insert(value int){
    newNode := &Node{value:value}
     if binaryTree.root == nil {
         binaryTree.root = newNode
     }else{
        tempRoots := make([] *Node,0)
        tempRoots  = append(tempRoots,binaryTree.root)
        ForLoop:
        for{
           currentRoot := tempRoots[0]
           tempRoots = tempRoots[1:]
           if value == currentRoot.value {
             break ForLoop
           }

           if currentRoot.left == nil{
             currentRoot.left = newNode
             break ForLoop
           }else{
             tempRoots = append(tempRoots,currentRoot.left)
           }

           if currentRoot.right == nil{
              currentRoot.right = newNode
              break ForLoop
           }else{
             tempRoots = append(tempRoots,currentRoot.right)
           }
        }
   }
}



func (binaryTree BinaryTree) PreOrder()(orderedSlice []int){

  nodeSlice := make([] *Node,0)
  nodeSlice = append(nodeSlice,binaryTree.root)
  for{
    tempNode :=  nodeSlice[len(nodeSlice)-1]
    orderedSlice = append(orderedSlice,tempNode.value)
    nodeSlice = nodeSlice[0:len(nodeSlice)-1]
    if tempNode.left!=nil{
         if tempNode.right!=nil{
           nodeSlice = append(nodeSlice,tempNode.right)
         }
          nodeSlice = append(nodeSlice,tempNode.left)
    }
   if len(nodeSlice)==0{
     break
   }
  }
   return
}



func (binaryTree BinaryTree) PostOrder()(orderedSlice []int){

  nodeSlice := make([] *Node,0)
  nodeSlice = append(nodeSlice,binaryTree.root)
  for{
    tempNode :=  nodeSlice[len(nodeSlice)-1]
    orderedSlice = append([]int{tempNode.value},orderedSlice...)
    nodeSlice = nodeSlice[0:len(nodeSlice)-1]
    if tempNode.right!=nil{
      nodeSlice = append(nodeSlice,tempNode.left)
      nodeSlice = append(nodeSlice,tempNode.right)
    }
    if tempNode.right==nil && tempNode.left!=nil{
      nodeSlice = append(nodeSlice,tempNode.left)
    }

   if len(nodeSlice)==0{
     break
   }
  }
   return
}


func (binaryTree BinaryTree) InOrder()(orderedSlice []int){
  nodeSlice := make([] *Node,0)
  currentNode := binaryTree.root
  for {
    if currentNode!=nil{
      nodeSlice = append(nodeSlice,currentNode)
      currentNode = currentNode.left
    }else if(len(nodeSlice)!=0){
      currentNode = nodeSlice[len(nodeSlice)-1]
      nodeSlice = nodeSlice[0:len(nodeSlice)-1]
      orderedSlice = append(orderedSlice,currentNode.value)
      currentNode = currentNode.right
    }else{
      break
    }
  }
   return
}


func RunBinaryTree(){
  var binaryTree = BinaryTree{}
  binaryTree.Insert(10)
  binaryTree.Insert(20)
  binaryTree.Insert(30)
  binaryTree.Insert(40)
  binaryTree.Insert(60)
  binaryTree.Insert(50)
  fmt.Println("Preorder Traversal : ",binaryTree.PreOrder())
  fmt.Println("Postorder Traversal : ",binaryTree.PostOrder())
  fmt.Println("Inorder Traversal : ",binaryTree.InOrder())
}
