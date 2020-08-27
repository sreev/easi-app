import { DateTime } from 'luxon';

import { LifecyclePhase } from 'types/estimatedLifecycle';

export type EstimatedLifecycleCostLines = {
  year1: LifecyclePhase[];
  year2: LifecyclePhase[];
  year3: LifecyclePhase[];
  year4: LifecyclePhase[];
  year5: LifecyclePhase[];
};

// Base Solution
export type BusinessCaseSolution = {
  title: string;
  summary: string;
  pros: string;
  cons: string;
  estimatedLifecycleCost: EstimatedLifecycleCostLines;
  costSavings: string;
};

// Proposed Solution
export type ProposedBusinessCaseSolution = BusinessCaseSolution & {
  acquisitionApproach: string;
  hosting: {
    type: string;
    location: string;
    cloudServiceType?: string;
  };
  hasUserInterface: string;
};

type BusinessCaseStatus = 'DRAFT' | 'SUBMITTED' | 'REVIEWED' | 'REJECTED';

export type GeneralRequestInfoForm = {
  requestName: string;
  requester: {
    name: string;
    phoneNumber: string;
  };
  businessOwner: {
    name: string;
  };
};

export type RequestDescriptionForm = {
  businessNeed: string;
  cmsBenefit: string;
  priorityAlignment: string;
  successIndicators: string;
};

export type AsIsSolutionForm = {
  asIsSolution: BusinessCaseSolution;
};

export type PreferredSolutionForm = {
  preferredSolution: ProposedBusinessCaseSolution;
};

export type AlternativeASolutionForm = {
  alternativeA: ProposedBusinessCaseSolution;
};

export type AlternativeBSolutionForm = {
  alternativeB: ProposedBusinessCaseSolution;
};

// Business Case Form Model
export type BusinessCaseModel = GeneralRequestInfoForm &
  RequestDescriptionForm &
  AsIsSolutionForm &
  PreferredSolutionForm &
  AlternativeASolutionForm &
  AlternativeBSolutionForm & {
    id?: string;
    euaUserId?: string;
    status: BusinessCaseStatus;
    systemIntakeId: string;
    initialSubmittedAt?: string;
    lastSubmittedAt?: string;
  };

export type BusinessCasesState = {
  businessCases: BusinessCaseModel[];
  isLoading: boolean | null;
  loadedTimestamp: DateTime | null;
  error: string | null;
};

// Redux store type for business case
export type BusinessCaseState = {
  form: BusinessCaseModel;
  isLoading: boolean | null;
  isSaving: boolean;
  isSubmitting: boolean;
  error: any;
};
