<template>
    <div :class="['node', nodeClass]">
      <div class="node-header" @click="toggleChildren">
        <span>{{ node.version }}</span> - <span>{{ node.node_type }}</span>
      </div>
      <div v-if="showChildren" class="node-children">
        <Node v-for="(child, index) in node.notes" :key="index" :node="child" />
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Node',
    props: {
      node: {
        type: Object,
        required: true
      }
    },
    data() {
      return {
        showChildren: false
      }
    },
    computed: {
      nodeClass() {
        return {
          created: 'node-created',
          anonymized: 'node-anonymized',
          signed: 'node-signed'
        }[this.node.node_type] || ''
      }
    },
    methods: {
      toggleChildren() {
        this.showChildren = !this.showChildren;
      }
    }
  }
  </script>
  
  <style scoped>
  .node {
    border: 1px solid #ccc;
    margin: 5px;
    padding: 10px;
    border-radius: 4px;
  }
  
  .node-header {
    font-weight: bold;
    cursor: pointer;
  }
  
  .node-children {
    margin-left: 20px;
  }
  
  /* Color scheme for different node types */
  .node-created {
    background-color: #e0f7fa; /* Light cyan */
    border-color: #00acc1; /* Dark cyan */
  }
  
  .node-anonymized {
    background-color: #f1f8e9; /* Light green */
    border-color: #9ccc65; /* Dark green */
  }
  
  .node-signed {
    background-color: #fce4ec; /* Light pink */
    border-color: #d81b60; /* Dark pink */
  }
  </style>
  