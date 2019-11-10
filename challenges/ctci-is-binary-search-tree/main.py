""" Node is defined as
class node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
    
    // Time: O(n) because every node is checked once
    // Space: O(n) in the worst case of a degenerate tree every node is in call stack
    //
    // The code structure with updating minimums and maximums in the recursion is necessary due to this case:
    //
    //         3
    //     2       6
    //   1  (4)  5   7
    //
    // The 4 there is larger than the 3 in the root node, but this memory is lost if we just check for immediately
    // connected nodes, which is what recursion gives us.
    //
    // The BST contract can be expressed in this one requirement (other than the fact that there's 2 children per node):
    // "Parent node value should be larger than ALL left children branch and smaller than ALL right children branch"
    // 
    // Thus, the main memory that needs to be passed downstream is the (min,max) boundaries of the recursed subtree.
"""
def checkBST(root):
    return doCheckBST(root, float('-inf'), float('+inf'))

def doCheckBST(root, minN, maxN):
    if root is None:
        return True
    return (root.data > minN and root.data < maxN and doCheckBST(root.left, minN, root.data) and doCheckBST(root.right, root.data, maxN))
