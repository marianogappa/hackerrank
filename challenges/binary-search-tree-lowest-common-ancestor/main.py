class Node:
    def __init__(self, info): 
        self.info = info  
        self.left = None  
        self.right = None 
        self.level = None 

    def __str__(self):
        return str(self.info) 

class BinarySearchTree:
    def __init__(self): 
        self.root = None

    def create(self, val):  
        if self.root == None:
            self.root = Node(val)
        else:
            current = self.root
         
            while True:
                if val < current.info:
                    if current.left:
                        current = current.left
                    else:
                        current.left = Node(val)
                        break
                elif val > current.info:
                    if current.right:
                        current = current.right
                    else:
                        current.right = Node(val)
                        break
                else:
                    break

# Enter your code here. Read input from STDIN. Print output to STDOUT
'''
class Node:
      def __init__(self,info): 
          self.info = info  
          self.left = None  
          self.right = None 
           

       // this is a node of the tree , which contains info as data, left , right
       //
       // Time: O(n) in the worst case of a degenerate tree. Avg O(logn)
       // Space: O(n) in the worst case of a degenerate tree; most nodes will be in the call stack. Avg O(logn)
       //
       // If both target nodes are larger  than the current node, then they must both be on the right branch.
       // If both target nodes are smaller than the current node, then they must both be on the left branch.
       // If one target node is equal to the current node, we're on the LCA (as long as we started from root).
       // If one target node is larger and the other smaller than current node, we're on the LCA (as long...).
'''
def lca(root, v1, v2):
  #Enter your code here
  if root.info < v1 and root.info < v2:
      return lca(root.right, v1, v2)
  if root.info > v1 and root.info > v2:
      return lca(root.left, v1, v2)
  return root

tree = BinarySearchTree()
t = int(input())

arr = list(map(int, input().split()))

for i in range(t):
    tree.create(arr[i])

v = list(map(int, input().split()))

ans = lca(tree.root, v[0], v[1])
print (ans.info)
