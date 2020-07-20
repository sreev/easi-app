import { DateTime } from 'luxon';

export type GovernanceCollaborationTeam = {
  collaborator: string;
  name: string;
};

type SystemIntakeStatus = 'DRAFT' | 'SUBMITTED' | 'REVIEWED' | 'REJECTED';

/**
 * Type for SystemIntakeForm
 *
 */
export type SystemIntakeForm = {
  id: string;
  euaUserID: string;
  requestName: string;
  status: SystemIntakeStatus;
  requester: {
    name: string;
    component: string;
  };
  businessOwner: {
    name: string;
    component: string;
  };
  productManager: {
    name: string;
    component: string;
  };
  isso: {
    isPresent: boolean | null;
    name: string;
  };
  governanceTeams: {
    isPresent: boolean | null;
    teams: GovernanceCollaborationTeam[];
  };
  fundingSource: {
    isFunded: boolean | null;
    fundingNumber: string;
  };
  businessNeed: string;
  businessSolution: string;
  currentStage: string;
  needsEaSupport: boolean | null;
  hasContract: string;
  businessCaseId?: string | null;
};

// Redux store type for a system intake
export type SystemIntakeState = {
  systemIntake: SystemIntakeForm;
  isLoading: boolean | null;
  isSaving: boolean;
  isSubmitting: boolean;
  error?: any;
};

// Redux store type for systems
export type SystemIntakesState = {
  systemIntakes: SystemIntakeForm[];
  isLoading: boolean | null;
  loadedTimestamp: DateTime | null;
  error: string | null;
};
