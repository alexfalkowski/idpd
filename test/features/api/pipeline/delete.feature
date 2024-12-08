Feature: Delete a pipeline
  The ability to delete pipelines from the system.

  Scenario: Delete a pipeline
    Given we create a simple pipeline
    When we delete the simple pipeline
    Then we should have deleted simple pipeline

  Scenario: Delete a nonexistent pipeline
    When we try to delete an nonexistent pipeline
    Then we should have a not found request

  Scenario: Delete pipeline by invalid id
    When we try to delete with an invalid id
    Then we should have bad request
