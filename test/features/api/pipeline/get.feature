Feature: Get a pipelines
  The ability to get pipelines from the system.

  Scenario: Get a simple pipeline
    Given we create a simple pipeline
    When we get the created simple pipeline
    Then we should have a simple pipeline

  Scenario: Get a nonexistent pipeline
    When we try to get an nonexistent pipeline
    Then we should have a not found request

  Scenario: Get a invalid pipeline
    When we get an invalid pipeline
    Then we should have bad request
