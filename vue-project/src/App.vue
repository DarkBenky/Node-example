<template>
  <div id="app">
    <div v-for="(node, index) in rootNodes" :key="index">
      <div class="head">Document:
        <Node :node="node" />
      </div>
    </div>
  </div>
</template>

<script>
import Node from './components/Node.vue'

export default {
  name: 'App',
  components: {
    Node
  },
  data() {
    return {
      rootNodes: [
        {
          "version": 1,
          "data": {
            "example": "Example"
          },
          "node_type": "created",
          "notes": [
            {
              "version": 1,
              "data": {
                "example": "Example"
              },
              "node_type": "anonymized",
              "notes": [
                {
                  "version": 1,
                  "data": {
                    "example": "Example"
                  },
                  "node_type": "signed",
                  "notes": []
                },
                {
                  "version": 1.1,
                  "data": {
                    "example": "Example"
                  },
                  "node_type": "signed",
                  "notes": []
                },
                {
                  "version": 1.2,
                  "data": {
                    "example": "Example"
                  },
                  "node_type": "signed",
                  "notes": []
                }
              ]
            },
            {
              "version": 1.1,
              "data": {
                "example": "Example"
              },
              "node_type": "anonymized",
              "notes": [
                {
                  "version": 1.1,
                  "data": {
                    "example": "Example"
                  },
                  "node_type": "signed",
                  "notes": []
                }
              ]
            },
            {
              "version": 1.2,
              "data": {
                "example": "Example"
              },
              "node_type": "anonymized",
              "notes": []
            },
            {
              "version": 1.3,
              "data": {
                "example": "Example"
              },
              "node_type": "anonymized",
              "notes": []
            }
          ]
        },
        {
          "version": 2,
          "data": {
            "example": "Another Example"
          },
          "node_type": "created",
          "notes": [
            {
              "version": 2,
              "data": {
                "example": "Another Example"
              },
              "node_type": "anonymized",
              "notes": [
                {
                  "version": 2,
                  "data": {
                    "example": "Another Example"
                  },
                  "node_type": "signed",
                  "notes": []
                }
              ]
            }
          ]
        }
      ]
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.head {
  font-size: 1.5em;
  background-color: #c6c6ec6b;
  color: black;
}
</style>
