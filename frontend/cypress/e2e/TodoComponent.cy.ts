/// <reference types="cypress" />

describe('template spec', () => {
  beforeEach(() => {
    cy.visit(Cypress.env('web_base_url'))
  })

  it('should add a new todo to the list', () => {
    cy.get('.input-wrap > input').type('SOS SOS')
    cy.get('button').click()

    cy.get('.input-wrap > input').type('SOS')
    cy.get('button').click()

    cy.get(':nth-child(1) > .task-wrapper > .task-title').should('have.text', 'SOS')
  })

  it('should mark as completed', () => {
    cy.get(':nth-child(1) > .task-wrapper > input').click()
    cy.get(':nth-child(1) > .task-wrapper > input').should('be.checked')

    cy.get(':nth-child(1) > .task-wrapper > .task-title').should(
      'have.css',
      'text-decoration-line',
      'line-through'
    )
    cy.get(':nth-child(1) > .task-wrapper > input').click()
  })

  it('should edit task', () => {
    cy.get(':nth-child(1) > .task-wrapper > :nth-child(4)').click()
    cy.get(':nth-child(1) > .task-wrapper > [type="text"]').type(' please')
    cy.get(':nth-child(1) > .task-wrapper > :nth-child(4)').click()

    cy.get(':nth-child(1) > .task-wrapper > .task-title').should('have.text', 'SOS please')
  })

  it('should delete task', () => {
    cy.get(':nth-child(1) > .task-wrapper > :nth-child(5)').click()

    cy.get(':nth-child(1) > .task-wrapper > .task-title').should('have.text', 'SOS SOS')
    cy.get(':nth-child(1) > .task-wrapper > :nth-child(5)').click()
  })
})
