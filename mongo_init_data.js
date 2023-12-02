

db = db.getSiblingDB('local')
db.createCollection('medical_articles')


db.medical_articles.insertMany([
    { title: 'test1', link: 'test_link1' },
    { title: 'test2', link: 'test_link2' },
    { title: 'test3', link: 'test_link3' },
    { title: 'test4', link: 'test_link4' }
  ]);