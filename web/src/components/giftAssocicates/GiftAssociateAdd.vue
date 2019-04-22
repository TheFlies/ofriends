<template>
  <el-dialog
    title="Assign Gift"
    :visible.sync="isVisibleAssign"
    width="30%"
    append-to-body
    :before-close="handleBackdropClick"
  >
    <div style="text-align: center">
      <el-transfer
        v-model="value"
        style="text-align: left; display: inline-block"
        filterable
        :filter-method="filterMethod"
        filter-placeholder="State Abbreviations"
        :titles="['Gifts', 'Assigned Gifts']"
        :data="data"
      />
      <div style="margin-top: 20px">
        <el-button
          type="primary"
          @click="submit()"
        >
          Save
        </el-button>
        <el-button @click="resetForm('giftAssociate')">
          Cancel
        </el-button>
      </div>
    </div>
  </el-dialog>
</template>

<script>
import { getAllGifts } from '@/api/gift'

export default {
  name: 'GiftAssociateAdd',

  props: {
    isVisibleAssign: { type: Boolean, default: false },
    assignedGifts: { type: Array, default: () => [] }
  },
  data() {
    return {
      data: [],
      value: []
    }
  },
  watch: {
    isVisibleAssign: function(val) {
      getAllGifts().then(resp => {
        this.value = []
        this.data = []
        if (resp.data != null) {
          resp.data.forEach((gift, index) => {
            this.data.push({
              label: gift.name,
              key: index,
              initial: gift.id
            })
            var position = this.assignedGifts.indexOf(gift.id)
            if (position >= 0) {
              this.value.push(index)
            }
          })
        }
      })
        .catch(err => {
          console.log(err)
        })
    }
  },
  methods: {
    filterMethod(query, item) {
      return item.initial.toLowerCase().indexOf(query.toLowerCase()) > -1
    },

    handleBackdropClick() {
      this.$emit('update:isVisibleAssign', false)
    },
    submit() {
      var giftAssociate = []
      this.value.forEach((val) => {
        giftAssociate.push(this.data[val])
      })
      this.$emit('update:isVisibleAssign', false)
      this.$emit('isGiftAssociateAdd', true, giftAssociate)
    },
    resetForm(formName) {
      this.$emit('update:isVisibleAssign', false)
    }
  }
}
</script>
