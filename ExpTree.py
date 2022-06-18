class Node:
    def __init__(self, data, leftChild=None, rightChild=None, parent=None):
        self.data = data
        self.left = leftChild
        self.right = rightChild
        self.parent = parent


class Stack:
    def __init__(self):
        self.data = []

    def push(self, data):
        self.data.append(data)

    def pop(self):
        return self.data.pop()


class ExpTree:
    def __init__(self, root=None):
        self.root = root

    def parseInfix(self, exp):
        stack = Stack()
        b = Node()  # branch
        for ch in exp:
            if ch == '(':  # Go left on (
                b = b.left
                continue
            if ch == ')':  # Go up on )
                b = b.parent
                continue
            if isOperator(ch):  # Set data on op
                continue
            if not isOperator(ch):  # Set leaf on operand
                stack.push(ch)

    def parsePrefix(self, exp):
        pass

    def parsePostfix(self, exp):
        pass

    def printInOrder(self):
        pass

    def printPreOrder(self):
        pass

    def printPostOrder(self):
        pass


def isOperator(ch):
    return ch in ['+', '-', '*', '/']


def main():
    print('Starting...')
    exp = '(2+3)'
    tree = ExpTree()
    tree.parseInOrder(exp)
    tree.printPreOrder()
    tree.printPostOrder()
    print('Done!')


if __name__ == '__main__':
    main()
