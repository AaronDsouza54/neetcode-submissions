class LRUCache {
    private HashMap<Integer, Integer> data;
    private LinkedList<Integer> usageTracking;
    int capacity;
    int size;

    public LRUCache(int capacity) {
        this.data = new HashMap<>();
        this.usageTracking = new LinkedList<>();
        this.size = 0;
        this.capacity = capacity;
    }
    
    public int get(int key) {
        var val = this.data.get(key);

        if (val == null) {
            return -1;
        }

        // We are making this data the most recently used
        this.usageTracking.remove((Integer) key);
        this.usageTracking.addFirst(key);

        return val;
    }
    
    public void put(int key, int value) {
        if (data.containsKey(key)) {
            this.data.put(key, value);
            this.usageTracking.remove((Integer) key);
            this.usageTracking.addFirst(key);
            return;
        }
        
        if (this.size == this.capacity) {
            int tail = this.usageTracking.removeLast();
            this.data.remove(tail);
            this.size--;
        }

        this.data.put(key, value);
        this.usageTracking.addFirst(key);
        this.size++;
    }
}
