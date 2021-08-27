import { Component, OnInit } from '@angular/core';
import { Store, select } from '@ngrx/store';
import { NgrxStateAtom } from 'app/ngrx.reducers';
import {
  GlobalDataFeedConfig
} from 'app/entities/destinations/destination.actions';
import {
  globalDataFeed
} from 'app/entities/destinations/destination.selectors';
import { Observable, Subject } from 'rxjs';
import { GlobalConfig } from 'app/entities/destinations/destination.model';
import { map } from 'rxjs/operators';

enum UrlTestState {
  Inactive,
  Loading,
  Success,
  Failure
}

@Component({
  selector: 'app-data-feed-config-details',
  templateUrl: './data-feed-config-details.component.html',
  styleUrls: ['./data-feed-config-details.component.scss']
})
export class DataFeedConfigDetailsComponent implements OnInit {
  private isDestroyed = new Subject<boolean>();
  public creatingDataFeed: boolean
  public hookStatus = UrlTestState.Inactive;
  public Fetchconfig$: Observable<GlobalConfig>;
  public config: GlobalConfig;
  getFirstFiveDataStatusCodes = true;
  getFirstFiveDataCidrFilters = true;
  constructor(
    private store: Store<NgrxStateAtom>,
  ) { }

  ngOnInit(): void {
    this.store.dispatch(new GlobalDataFeedConfig({}))
    this.Fetchconfig$ = this.store.pipe(
                    select(globalDataFeed)).pipe(
                      map((res: GlobalConfig) =>  res));
   this.Fetchconfig$.subscribe((res: GlobalConfig) => {
     this.config = res
   })
  }
  ngOnDestroy(): void {
    this.isDestroyed.next(true);
    this.isDestroyed.complete();
  }
  covertFeedInterval(feedInterval:string):string{
    if(feedInterval != null){
      let SplitFeedInterval = feedInterval.split('')
      if(SplitFeedInterval.includes('h')){
        SplitFeedInterval.splice(SplitFeedInterval.indexOf('h'), 1, " Hour ");
      }
      if(SplitFeedInterval.includes('m')){
        SplitFeedInterval.splice(SplitFeedInterval.indexOf('m'), 1, " Minute ");
      }
      if(SplitFeedInterval.includes('s')){
        SplitFeedInterval.splice(SplitFeedInterval.indexOf('s'), 1, " Seconds ");
      }
      return SplitFeedInterval.join('')
    }
  }

  getAlldata(flag:string):void{
    if(flag == 'StatusCodes'){
      this.getFirstFiveDataStatusCodes = false
    } else if(flag == 'DataCidrFilters') {
      this.getFirstFiveDataCidrFilters = false
    }
  }
  getData(arrValue?:any,getFirstFiveData?:boolean):any{
    if(arrValue != null){
      if(getFirstFiveData){
        if(arrValue.length>=5){
          return arrValue.slice(0,5)
        }else{
          return arrValue.slice(0,arrValue.length)
        }
      } else {
        return arrValue
      }
    }
  }
}
