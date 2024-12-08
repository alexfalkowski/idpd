Feature: Trigger a pipeline
  The ability to trigger pipelines and get the output.

  Scenario: Trigger a pipeline
    Given we create a simple pipeline
    When we trigger the simple pipeline
    Then we should receive the output of simple pipeline

  Scenario: Trigger an invalid pipeline
    Given we create an invalid pipeline
    When we trigger the simple pipeline
    Then we should receive that the pipeline contained errors

  Scenario: Trigger a nonexistent pipeline
    When we try to trigger an nonexistent pipeline
    Then we should have a not found request

  Scenario: Trigger pipeline by invalid id
    When we try to trigger with an invalid id
    Then we should have bad request
