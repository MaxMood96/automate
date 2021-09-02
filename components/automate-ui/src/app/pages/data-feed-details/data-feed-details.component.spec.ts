import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { waitForAsync, ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { ReactiveFormsModule, FormsModule, FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Store, StoreModule } from '@ngrx/store';
import { ngrxReducers, NgrxStateAtom, runtimeChecks } from '../../../app/ngrx.reducers';
import { DataFeedDetailsComponent } from './data-feed-details.component';
import { FeatureFlagsService } from '../../../app/services/feature-flags/feature-flags.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { Destination } from 'app/entities/destinations/destination.model';
import { Regex } from 'app/helpers/auth/regex';
import { By } from '@angular/platform-browser';

describe('DataFeedDetailsComponent', () => {
    let component: DataFeedDetailsComponent;
    let fixture: ComponentFixture<DataFeedDetailsComponent>;
    let updateForm: FormGroup;
    let store: Store<NgrxStateAtom>;

    beforeEach(waitForAsync(() => {
      TestBed.configureTestingModule({
        declarations: [
          DataFeedDetailsComponent
        ],
        providers: [
          FeatureFlagsService
        ],
        imports: [
          FormsModule,
          ReactiveFormsModule,
          RouterTestingModule,
          BrowserAnimationsModule,
          StoreModule.forRoot(ngrxReducers, { runtimeChecks })
        ],
        schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
      })
      .compileComponents();
    }));

    beforeEach(() => {
      fixture = TestBed.createComponent(DataFeedDetailsComponent);
      store = TestBed.inject(Store);
      spyOn(store, 'dispatch').and.callThrough();
      component = fixture.componentInstance;
      component.updateForm = new FormBuilder().group({
        name: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]],
        // Note that URL here may be FQDN -or- IP!
        url: ['', [Validators.required, Validators.pattern(Regex.patterns.VALID_FQDN)]],
        tokenType: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]],
        token: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]],
        username: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]],
        password: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]]
      });
      updateForm = component.updateForm;
      fixture.detectChanges();
    });

    it('should create', () => {
      expect(component).toBeTruthy()
    });
    
    describe('Data Feed overview', () => {
      const destination = <Destination> {
        id: '1',
        name: 'new data feed',
        secret: 'testSecret',
        url: 'http://foo.com'
      };

      const updatedDestination = <Destination> {
        id: '1',
        name: 'updated data feed',
        secret: 'testSecret',
        url: 'http://foo.com/serviceNow"'
      };

      it('check is value is set in form', () => {
        component.updateForm.controls['name'].setValue(destination.name);
        component.updateForm.controls['url'].setValue(destination.url);
        expect(component.updateForm.controls['name'].value).toEqual(destination.name);
        expect(component.updateForm.controls['url'].value).toEqual(destination.url);
      });

      it('check enableDestination func', () => {
        component.destination = destination
        component.enableDestination(true)
        expect(component.destination.enable).toBeTrue
        component.enableDestination(false)
        expect(component.destination.enable).toBeFalse
      })

      it('check saveDataFeed func', async () => {
        component.destination = destination
        component.updateForm.controls['name'].setValue(updatedDestination.name);
        component.updateForm.controls['url'].setValue(updatedDestination.url);
        let buttonElement = fixture.debugElement.query(By.css('.save-connection'));
        buttonElement.triggerEventHandler('click', null);        
        expect(component.destination.name).toEqual(updatedDestination.name)
        expect(component.destination.url).toEqual(updatedDestination.url)

      })

      it('check sendTestForDataFeedUrl func', () => {
        component.destination = destination
        component.testInProgress = true
        let buttonElement = fixture.debugElement.query(By.css('.test-connection'));
        buttonElement.triggerEventHandler('click', null);
        expect(component.testInProgress).toBeFalse()

      })
      it('check deleteModalVisible func', () => {
        component.deleteModalVisible = true
        component.closeDeleteModal
        expect(component.deleteModalVisible).toBeFalse
      })
      it('check openDeleteModal func', () => {
        component.deleteModalVisible = false
        component.openDeleteModal
        expect(component.deleteModalVisible).toBeFalse
      })
    })
  });
