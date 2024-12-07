Feature: Pipelines
  A continuous integration and continuous deployment (CI/CD) pipeline is a series of steps that must be performed in order to deliver a new version of software.

  Scenario: Create a pipeline
    When we create a simple pipeline
    Then we should have a created pipeline

  Scenario Outline: Create invalid pipeline
    When we try to create a invalid pipeline of kind "<kind>"
    Then we should have bad request

    Examples:
      | kind                  |
      | invalid pipeline name |
      | missing jobs          |
      | invalid job name      |
      | missing steps         |

  Scenario: Unauthorized access
    When we try to create a pipeline with being authorized
    Then we should have a unauthorized request

  Scenario: Get a pipeline
    Given we create a simple pipeline
    When we get the created simple pipeline
    Then we should have a simple pipeline

  Scenario: Get a nonexistent pipeline
    When we try to get an nonexistent pipeline
    Then we should have a not found request

  Scenario: Get a invalid pipeline
    When we get an invalid pipeline
    Then we should have bad request

  Scenario: Update a pipeline
    Given we create a simple pipeline
    When we get the update the simple pipeline
    Then we should have an updated simple pipeline

  Scenario: Update a nonexistent pipeline
    When we try to update an nonexistent pipeline
    Then we should have a not found request

  Scenario: Update an invalid pipeline
    When we try to update an invalid pipeline
    Then we should have bad request

  Scenario: Update pipeline by invalid id
    When we try to update with an invalid id
    Then we should have bad request
