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
import { EntityStatus } from 'app/entities/entities';

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
  public configNotFound: boolean = true;
  getFirstFiveDataStatusCodesShow:boolean;
  getStatusCodes: number[] = [];
  getCidrFilters:string[] = [];
  getFirstFiveDataCidrFiltersShow:boolean;
  constructor(
    private store: Store<NgrxStateAtom>,
  ) { }

  ngOnInit(): void {
    this.store.dispatch(new GlobalDataFeedConfig({}))
      this.getGlobalDataFeedConfig()
}

  public getGlobalDataFeedConfig(){
    this.store.pipe(
      select(globalDataFeed)).pipe(
        map((res: any) =>  res))
        .subscribe((res: any) => {
          if(res.globalConfigStatus ==  EntityStatus.loadingSuccess){
            this.configNotFound = false
            this.config = res.globalConfig
            this.getStatusCodes = this.getData(this.config?.accepted_status_codes,true)
            this.getCidrFilters = this.getData(this.config?.cidr_filter,true)
            this.getFirstFiveDataStatusCodesShow = this.config?.accepted_status_codes.length == this.getStatusCodes.length
            this.getFirstFiveDataCidrFiltersShow = this.config?.cidr_filter.length == this.getCidrFilters.length
          }
        })
  }
  ngOnDestroy(): void {
    this.isDestroyed.next(true);
    this.isDestroyed.complete();
  }
  covertFeedInterval(feedInterval:string):string{
    let output: string[] = []
    if(feedInterval != null){
      let SplitFeedInterval = feedInterval.split('')
      if(SplitFeedInterval.includes('h')){
        SplitFeedInterval.splice(SplitFeedInterval.indexOf('h'), 1, " Hour ,");
      }
      if(SplitFeedInterval.includes('m')){
        SplitFeedInterval.splice(SplitFeedInterval.indexOf('m'), 1, " Minute ,");
      }
      if(SplitFeedInterval.includes('s')){
        SplitFeedInterval.splice(SplitFeedInterval.indexOf('s'), 1, " Seconds ,");
      }
      let SplitIntoTime = SplitFeedInterval.join('')
      SplitIntoTime.split(',').forEach((v) => {
        if (v.includes('Hour')) {
          if(parseInt(v) != 0){
            output.push(v)
          }
        }
        if (v.includes('Minute')) {
          if(parseInt(v)!= 0){
            output.push(v)
          }
        }
        if (v.includes('Seconds')) {
          if(parseInt(v)!= 0){
            output.push(v)
          }
        }
      })      
      return output.join('')
    }
  }

  getAlldata(flag:string):void{
    if(flag == 'StatusCodes'){
      this.getFirstFiveDataStatusCodesShow = false;
      this.getStatusCodes = this.getData(this.config.accepted_status_codes,false)
      this.getFirstFiveDataStatusCodesShow = this.config?.accepted_status_codes.length == this.getStatusCodes.length
    } else if(flag == 'DataCidrFilters') {
      this.getFirstFiveDataCidrFiltersShow = false;
      this.getCidrFilters = this.getData(this.config.cidr_filter,false)
      this.getFirstFiveDataCidrFiltersShow = this.config?.cidr_filter.length == this.getCidrFilters.length
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
    return []
  }
}

