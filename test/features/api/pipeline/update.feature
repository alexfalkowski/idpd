Feature: Update a pipeline
  The ability to update pipelines in the system.

  Scenario: Update a pipeline
    Given we create a simple pipeline
    When we update the simple pipeline
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
