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

       // Time: O(n) in the worst case of a degenerate tree: we must traverse all nodes. Hopefully O(logn) avg.
       // Space: O(n) in the worst case of a degenerate tree: all nodes in the call stack. Hopefully O(logn) avg.
       //
       // Height in terms of nodes rather than edges is trivial with recursion: this is what nodeBaseHeight does.
       // Edge-based height is node-based height minus one, except when the root node is null.
'''
def height(root):
    if root == None:
        return 0
    return nodeBasedHeight(root) - 1

def nodeBasedHeight(root):
    if root == None:
        return 0
    return 1 + max(nodeBasedHeight(root.left), nodeBasedHeight(root.right))



tree = BinarySearchTree()
t = int(input())

arr = list(map(int, input().split()))

for i in range(t):
    tree.create(arr[i])

print(height(tree.root))
