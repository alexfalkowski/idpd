Feature: Create a pipeline
  The ability to add pipeline to the system.

  Scenario: Create a simple pipeline
    When we create a simple pipeline
    Then we should have a created pipeline

  Scenario Outline: Create an invalid pipeline
    When we try to create a invalid pipeline of kind "<kind>"
    Then we should have bad request

    Examples:
      | kind                  |
      | invalid pipeline name |
      | missing jobs          |
      | invalid job name      |
      | missing steps         |

  Scenario: Create pipeline with unauthorized access
    When we try to create a pipeline with being authorized
    Then we should have a unauthorized request

  Scenario: Create pipeline with a bad payload
    When we try to create a pipeline with a bad payload
    Then we should have bad request

  Scenario: Create pipeline with a empty payload
    When we try to create a pipeline with an empty payload
    Then we should have bad request
