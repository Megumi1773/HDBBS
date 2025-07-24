import Mock from 'mockjs'

// 定义接口
Mock.mock('/api/categories', 'get', {
  code: 200,
  data: [{
      ID: 1,
      CategoryName: '学习交流',
      CategoryIcon: '1'
    },
    {
      ID: 2,
      CategoryName: '求职招聘',
      CategoryIcon: '2'
    }
  ],
  msg: '获取成功'
})