DO $$ BEGIN
		CREATE TYPE vendor AS ENUM ('naver', 'kakao');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;