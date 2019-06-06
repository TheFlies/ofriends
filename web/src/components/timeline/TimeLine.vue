<template>
  <el-container style="background: #ecf5ff;">
    <el-main>
      <el-table
        v-loading="loading"
        :data="tableData.filter(data => !search || data.Customer.name.toLowerCase().includes(search.toLowerCase()))"
        style="width: 100%; margin:auto"
      >
        <el-table-column type="index" :index="indexMethod" />
        <el-table-column label="Arrived Time" width="180" sortable prop="Visit.arrivedTime">
          <template slot-scope="scope">
            {{ getHumanDate(scope.row.Visit.arrivedTime) }}
          </template>
        </el-table-column>
        <el-table-column label="Departure Time" width="180" sortable prop="Visit.departureTime">
          <template slot-scope="scope">
            {{ getHumanDate(scope.row.Visit.departureTime) }}
          </template>
        </el-table-column>
        <el-table-column label="Visit Name" prop="Visit.name" sortable width="180">
          <template slot-scope="scope">
            <el-popover
              placement="right"
              width="400"
              trigger="click"
            >
              <el-form
                ref="scope.row.Visit"
                :model="scope.row.Visit"
                label-width="120px"
              >
                <h3 class="header-popover">
                  Visit Info
                </h3>
                <el-form-item label="Name" class="label">
                  <span>{{ scope.row.Visit.name }}</span>
                </el-form-item>
                <el-form-item label="Arrived Time" class="label">
                  <el-tag v-for="tag in scope.row.Visit.lab" :key="tag" size="small" type="success" :disable-transitions="false">
                    {{ tag }}
                  </el-tag>
                </el-form-item>
                <el-form-item label="Arrived Time" class="label">
                  <span>{{ getHumanDate(scope.row.Visit.arrivedTime) }}</span>
                </el-form-item>
                <el-form-item label="Departure Time" class="label">
                  <span>{{ getHumanDate(scope.row.Visit.departureTime) }}</span>
                </el-form-item>
                <el-form-item label="Accommodation" class="label">
                  <span>{{ scope.row.Visit.accommodation }}</span>
                </el-form-item>
                <el-form-item label="Pickup" class="label">
                  <span>{{ scope.row.Visit.pickup }}</span>
                </el-form-item>
              </el-form>
              <el-button slot="reference">
                {{ scope.row.Visit.name }}
              </el-button>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column label="Visit Lab" width="130" prop="Visit.lab">
          <template slot-scope="scope">
            <el-tag v-for="tag in scope.row.Visit.lab" :key="tag" size="small" type="success" :disable-transitions="false">
              {{ tag }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Customer Name" width="180" sortable>
          <template slot-scope="scope">
            <el-popover
              placement="right"
              width="600"
              trigger="click"
            >
              <el-form
                ref="scope.row.Customer"
                :model="scope.row.Customer"
                label-width="120px"
              >
                <h3 class="header-popover">
                  Customer Info
                </h3>
                <el-row :gutter="20">
                  <el-col :span="12">
                    <div class="grid-content bg-purple">
                      <el-form-item label="Name" class="label">
                        <span>{{ scope.row.Customer.name }}</span>
                      </el-form-item>
                      <el-form-item label="Title" class="label">
                        <span>{{ scope.row.Customer.title }}</span>
                      </el-form-item>
                      <el-form-item label="Position" class="label">
                        <span>{{ scope.row.Customer.position }}</span>
                      </el-form-item>
                      <el-form-item label="Project" class="label">
                        <span>{{ scope.row.Customer.project }}</span>
                      </el-form-item>
                      <el-form-item label="Age" class="label">
                        <span>{{ scope.row.Customer.age }}</span>
                      </el-form-item>
                      <el-form-item label="Company" class="label">
                        <span>{{ scope.row.Customer.company }}</span>
                      </el-form-item>
                      <el-form-item label="Country" class="label">
                        <span>{{ scope.row.Customer.country }}</span>
                      </el-form-item>
                      <el-form-item label="City" class="label">
                        <span>{{ scope.row.Customer.city }}</span>
                      </el-form-item>
                    </div>
                  </el-col>
                  <el-col :span="12">
                    <div class="grid-content bg-purple">
                      <el-form-item label="Passport info" class="label">
                        <span>{{ scope.row.Customer.passportInfo }}</span>
                      </el-form-item>
                      <el-form-item label="Food Note" class="label">
                        <span>{{ scope.row.Customer.foodNote }}</span>
                      </el-form-item>
                      <el-form-item label="Family Note" class="label">
                        <span>{{ scope.row.Customer.familyNote }}</span>
                      </el-form-item>
                      <el-form-item label="Next visit Note" class="label">
                        <span>{{ scope.row.Customer.nextVisitNote }}</span>
                      </el-form-item>
                    </div>
                  </el-col>
                </el-row>
              </el-form>
              <el-button slot="reference">
                {{ scope.row.Customer.name }}
              </el-button>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column label="Position" prop="Customer.position" width="100" />
        <el-table-column label="Project" sortable width="150" prop="Customer.project" />
        <el-table-column label="Gift" width="130" sortable>
          <template slot-scope="scope">
            <el-popover
              v-for="gift in scope.row.Gifts"
              :key="gift.id"
              placement="left"
              width="400"
              trigger="click"
            >
              <el-form
                ref="gift"
                :model="gift"
                label-width="80px"
              >
                <h3 class="header-popover">
                  Gift Info
                </h3>
                <el-form-item label="Name" class="label">
                  <span>{{ gift.name }}</span>
                </el-form-item>
                <el-form-item label="Idea" class="label">
                  <span>{{ gift.idea }}</span>
                </el-form-item>
                <el-form-item label="Size" class="label">
                  <span>{{ gift.size }}</span>
                </el-form-item>
                <el-form-item label="Price" class="label">
                  <span>{{ gift.price }}</span>
                </el-form-item>
                <el-form-item label="Link" class="label">
                  <a :href="`${gift.link}`" target="_blank">
                    {{ gift.link }}
                  </a>
                </el-form-item>
                <el-form-item label="Description" class="label">
                  <span>{{ gift.description }}</span>
                </el-form-item>
              </el-form>
              <el-button slot="reference">
                {{ gift.name }}
              </el-button>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column label="Activity" sortable prop="priority">
          <template slot-scope="scope">
            <el-popover
              v-for="activity in scope.row.Activities"
              :key="activity.id"
              placement="left"
              width="400"
              trigger="click"
            >
              <el-form
                ref="activity"
                :model="activity"
                label-width="80px"
              >
                <h3 class="header-popover">
                  Activity Info
                </h3>
                <el-form-item label="Name" class="label">
                  <span>{{ activity.name }}</span>
                </el-form-item>
                <el-form-item label="Detail" class="label">
                  <span>{{ activity.detail }}</span>
                </el-form-item>
              </el-form>
              <el-button slot="reference">
                {{ activity.name }}
              </el-button>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column align="right">
          <template slot="header" slot-scope="scope">
            <el-input
              v-model="search"
              size="mini"
              placeholder="Customer name"
            />
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>
<script>
import { getTimelineByDay } from '@/api/timeline'
import { indexMethod, getHumanDate } from '@/utils/convert'

export default {
  name: 'Timeline',
  components: {
  },
  data() {
    return {
      tableData: [],
      search: '',
      loading: true
    }
  },
  mounted() {
    getTimelineByDay(Date.now())
      .then(resp => {
        if (resp.data != null) {
          this.tableData = resp.data
        }
        this.loading = false
      })
      .catch(err => {
        console.log(err)
      })
  },
  methods: {
    indexMethod,
    getHumanDate
  }
}
</script>
<style scoped lang="stylus">
.label {
  color: #259dd8;
}
.header-popover{
  text-align: center;
  background-color: #259dd8;
  color: #fff;
  padding: 10px;
}
</style>
