class Node:
  def __init__(self, data=None, next=None, prev=None):
    self.next: Node = next
    self.prev: Node = prev
    self.data = data

class DoublyLinkedList:
  def __init__(self, head: Node = None):
    self.head = head
    self.last: Node = None
    if head is not None:
      node = head
      while node:
        if node.next is None:
          self.last = node
        node = node.next

  def pushData(self, newData):
    self.pushNode(Node(newData))

  def pushNode(self, newNode: Node):
    newNode.prev = self.last
    if self.head is None:
      self.head = newNode
    if self.last is not None:
      self.last.next = newNode
    self.last = newNode

  def insertData(self, newData):  # Ordered
    self.insertNode(Node(newData))

  def insertNode(self, newNode: Node):  # Ordered
    if self.head is None:
      return self.pushNode(newNode)
    node = self.head
    if node.data > newNode.data:
        newNode.next = node
        newNode.prev = node.prev
        node.prev = newNode
        self.head = newNode
        return
    while node:
      if node.next is None:  # if last
        return self.insertNodeAfter(node, newNode)
      if newNode.data < node.next.data:
        return self.insertDataAfter(node, newNode)
      node = node.next

  def removeData(self, data):
    node = self.head
    while node:
      if node.data == data:
        self.removeNode(node)
      node = node.next

  def removeNode(self, node: Node):
    if node.prev:
      node.prev.next = node.next
    if node.next:
      node.next.prev = node.prev
    node, node.prev, node.next = None

  def insertDataAfter(self, prevNode: Node, newData):
    self.insertNodeAfter(Node(newData, prevNode.next, prevNode))

  def insertNodeAfter(self, prevNode: Node, newNode):
    if newNode.next:
      prevNode.next.prev = newNode
    prevNode.next = newNode

  def __str__(self):
    node = self.head
    string = '<=> '
    while node:
      string += str(node.data) + ' <=> '
      node = node.next
    # string += ']'
    return string


def main():
  print('Hello, World!')
  l = DoublyLinkedList()
  print(l)
  l.insertData(2)
  print(l)
  l.insertData(4)
  print(l)
  l.insertData(1)
  print(l)
  l.insertData(12)
  print(l)
  l.insertData(0)


if __name__ == '__main__':
  main()
