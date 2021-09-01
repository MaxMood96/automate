describe('chef datafeed', () => {

    const token = 'behwveh3238238=';
    const DestinationID = 'chef-server-dev-test';
    const reusableDate = Date.now();
    let adminIdToken = '';
    let name = 'cytest' + reusableDate
    const url = "https://localhost/api/x_chef_automate/asset"
    const ServiceNow = "Service Now"
    const Webhook = "Webhook"
    before(() => {
        cy.adminLogin('/').then(() => {
            const admin = JSON.parse(<string>localStorage.getItem('chef-automate-user'));
            adminIdToken = admin.id_token;
            cy.get('app-welcome-modal').invoke('hide');
            cy.restoreStorage();
            cy.get('body').type('feat');
            cy.get('.title').contains('Chef Automate Data Feed').parent().parent()
              .find('.onoffswitch').click();
            cy.get('chef-button').contains('Close').click();
            cy.reload();
            cy.contains('know').click();
        });
        cy.restoreStorage();
      });

    beforeEach(() => {
        cy.restoreStorage();
      });
  
    afterEach(() => {
        cy.saveStorage();
    });
    describe ('chef data feed page for Service Now', () => {
        before(() => {
            cy.request({
                auth: { bearer: adminIdToken },
                method: "POST",
                url: "/api/v0/secrets",
                body: {
                  name: "anyname",
                  type: "data_feed",
                  data: [
                      {
                          key: "headers",
                          value: "{\"Authorization\":\"Basic service-now-token\"}"
                      },
                      {
                          key: "auth_type",
                          value: ""
                      }
                  ]
                }
            }).then((resp) => {
                  expect(resp.status).to.equal(200);
                  expect(resp.body).to.property('id') 
                  cy.request({
                      auth: { bearer: adminIdToken },
                      method: "POST",
                      url: "/api/v0/datafeed/destination",
                      body: {
                          url:url,
                          name:name,
                          secret:resp.body.id,
                          services:ServiceNow,
                          integration_types:Webhook
                      }
                  }).then((resp) => {
                    expect(resp.status).to.equal(200);
                    expect(resp.body).to.property('id') 
                    cy.visit(`/settings/data-feeds/${resp.body.id}`);
                  })
            })
        })
        it('check if clicking on disable', () => {
            cy.get('[data-cy=disable-btn]').click();
            cy.get('app-notification.info').should('be.visible');
            cy.get('app-notification.info p').contains("Disabled").should('exist');
            cy.get('app-notification.info chef-icon').click();
        })
        it('check if clicking on enable', () => {
            cy.get('[data-cy=enable-btn]').click();
            cy.get('app-notification.info').should('be.visible');
            cy.get('app-notification.info p').contains("Enabled").should('exist');
            cy.get('app-notification.info chef-icon').click();
        })
        it('check if clicking on save after changing name and url ', () => {
            cy.get('[data-cy="name-input"]').type("-1")
            cy.get('[data-cy="url-input"]').type("/v1")
            cy.get('[data-cy="save-connection"]').click();
            cy.get('app-notification.info').should('be.visible');
            cy.get('app-notification.info p').contains(name+"-1").should('exist');
            cy.get('app-notification.info chef-icon').click();
        })
        it('check if clicking on test connection', () => {
            cy.get('[data-cy="name-input"]').type("-1")
            cy.get('[data-cy="url-input"]').type("/v1")
            cy.get('[data-cy="save-connection"]').click();
            cy.get('app-notification.error').should('be.visible');
            cy.get('app-notification.info chef-icon').click();
        })
        it('check if clicking on Delete', () => {
            cy.get('[data-cy=delete-btn]').click();
            cy.get('app-delete-object-modal').find('button').contains('Delete Data Feed')
            .click({force: true});
            cy.get('chef-tbody').contains(name).should('not.exist');
        })
    })
    
})