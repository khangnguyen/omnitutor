``Database Schema``
===================

The schema listing given here is not intended to be comprehensive. It is intended to give a general structure and address some specific decisions, but other tables may be added and those given here may be expanded.

.. contents::

``Accounts``
------------

``user_account``
````````````````

============================= ============== ====================================
Column                        Type           Notes
============================= ============== ====================================
``id``                        serial         Primary Key
``nickname``                  varchar        Display name
``email``                     varchar        Unique
``username``                  varchar        Unique
``password``                  varchar        BCrypt hash of password
``gender``                    varchar        ``male``, ``female``, ``other``
``avatar_filename``           text
``inviter_id``                int            References ``user_account(id)``
``created_at``                timestamp
``updated_at``                timestamp
``last_login_at``             timestamp
``is_archived``               boolean        Default False
``locale``                    varchar        Locale setting of the user. Default to ``en``
============================= ============== ====================================

``verification_token``
``````````````````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``code``                uuid         Primary key
``user_account_id``     int          References ``user_account(id)``
``validation_time``     timestamp    One hour from created time
``is_valid``            bool         Default true
======================= ============ ====================================

``user_login_oauth``
````````````````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``id``                  serial
``user_account_id``     int          References ``user_account(id)``
``service``             varchar      Third-party OAuth service used
``identity``            varchar      Identity in service
``access_token``        varchar      Issued by service provider
======================= ============ ====================================

**Constraint**: Unique on (``user_account_id``, ``service``).

``Course Structure``
--------------------

``institute``
`````````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``id``                  serial
``user_account_id``     int          References ``user_account(id)``
``name``                varchar      e.g. 'The Tutor House'
``description``         text
``logo_filename``       varchar      
``cover_filename``      varchar
``slider_images``       JSON         The list of slider images in game detail: { 'images': [ {'filename', 'index'}, ... ] }
``created_at``          timestamp
``updated_at``          timestamp
``is_active``           boolean      Default False
``is_archived``         boolean      Default False
======================= ============ ====================================

``institude_student``
`````````````````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``institute_id``        int          References ``institute(id)``
``student_id``          int          References ``user_account(id)``
``joined_on``           timestamp
``source``              varchar      e.g. Facebook, Website, Import
======================= ============ ====================================

``course``
``````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``id``                  serial
``institute_id``        int          References ``institute(id)``
``name``                varchar      e.g. 'IELTS 101'      
``description``         text
``logo_filename``       varchar      
``created_at``          timestamp
``updated_at``          timestamp
``is_active``           boolean      Default False
``is_archived``         boolean      Default False
======================= ============ ====================================

``skill``
`````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``id``                  serial
``institute_id``        int          References ``institute(id)``
``name``                varchar      e.g. 'Analyze graphs'
``description``         text
``content``             text
``homework``            text
``created_at``          timestamp
``updated_at``          timestamp
======================= ============ ====================================

``group``
`````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``id``                  serial
``name``                varchar      e.g. 'Morning Mon-Wed-Fri'
``course_id``           int          References ``course(id)``
``description``         text
``created_at``          timestamp
``updated_at``          timestamp
``is_active``           boolean      Default False
``is_archived``         boolean      Default False
======================= ============ ====================================

``group_student``
`````````````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``group_id``            int          References ``group(id)``
``student_id``          int          References ``user_account(id)``
======================= ============ ====================================

**Constraint**: Unique on (``group_id``, ``student_id``).

``class``
`````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``id``                  serial
``group_id``            int          References ``group(id)``
``skill_id``            int          References ``skill(id)``
``start_time``          timestamp
``end_time``            timestamp
======================= ============ ====================================


``class_student``
`````````````````

======================= ============ ====================================
Column                  Type         Notes
======================= ============ ====================================
``class_id``            int          References ``class(id)``
``student_id``          int          References ``user_account(id)``
``attented``            boolean      Default False
======================= ============ ====================================

**Constraint**: Unique on (``class_id``, ``student_id``).
