// Copyright (©) 2024 - Shivam Kumar Jha - All Rights Reserved, Proprietary and confidential
// Unauthorised copying of this file, via any medium is strictly prohibited

import { knex, randomBool, randomInt } from './util.js'
import { fakerEN_IN as faker } from '@faker-js/faker'

// Generate Classes
export const generateClass = async () => {
  const classes = []
  for (let index = 1; index < 13; index++) {
    classes.push({ name: `${index}` })
  }
  await knex('class').insert(classes)
}

export const generateSection = async () => {
  // Generate Sections
  const classes = await knex('class').select('*')
  for (const cls of classes) {
    const sectionsCount = randomInt(1, 5)
    const sections = []
    for (let index = 0; index < sectionsCount; index++) {
      sections.push({
        name: String.fromCharCode(65 + index),
        classid: cls.id,
      })
    }
    await knex('section').insert(sections)
  }
}

// Generate Students for every section
export const generateUser = async () => {
  const allUsers = [{ firstname: 'Shivam', lastname: 'Kumar Jha', email: 'shivamkj360@gmail.com' }]
  // female students
  const femaleCount = randomInt(17, 23)
  for (let index = 0; index < 1500; index++) {
    const user = {
      firstname: faker.person.firstName(),
      lastname: faker.person.lastName(),
    }
    const identity = randomInt(0, 2)
    user.email = identity == 0 || identity == 2 ? faker.internet.email({ ...user }) : null
    user.mobilenum = identity == 1 || identity == 2 ? faker.string.numeric(10) : null
    user.profilepic = randomBool(0.8) ? faker.image.avatar() : null
    allUsers.push(user)
  }
  await knex('users').insert(allUsers).onConflict(['email', 'mobilenum']).ignore()
}
